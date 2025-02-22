package pkg

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type DataFile struct {
	Header Header `xml:"header"`
	Games  []Game `xml:"game"`
}

type Header struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

type Game struct {
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description"`
	Releases    []Release `xml:"release"`
	Roms        []Rom     `xml:"rom"`
}

type Release struct {
	Name   string `xml:"name,attr"`
	Region string `xml:"region,attr"`
}

type Rom struct {
	Name   string `xml:"name,attr"`
	Size   string `xml:"size,attr"`
	CRC    string `xml:"crc,attr"`
	MD5    string `xml:"md5,attr"`
	SHA1   string `xml:"sha1,attr"`
	Status string `xml:"status,attr"`
}

type FileHashes struct {
	MD5  string
	SHA1 string
}

type RomInfo struct {
	GameName string
	RomName  string
}

// Regular expression to remove leading digits and optional separators
var reLeadingNumbers = regexp.MustCompile(`^[0-9]+(\s*-\s*)?`)

// ProcessRoms receives directory paths and file suffix as inputs, performs the scanning, renaming, and logging operations.
func ProcessRoms(dirPath, datPath, fileSuffix string) error {
	// Create/append a log file in the target directory
	logFilePath := filepath.Join(dirPath, "resultado_log.txt")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening/creating log file: %v", err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "", log.LstdFlags)

	// 1) Load the .dat file
	data, err := loadDatFile(datPath)
	if err != nil {
		logger.Printf("Error loading .dat file: %v\n", err)
		return err
	}

	// 2) Build ROM hash lookup maps
	md5Map, sha1Map := buildRomMaps(data)
	logger.Printf("Loaded .dat with %d games.", len(data.Games))

	// 3) Walk through the directory
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			logger.Printf("Error accessing '%s': %v", path, walkErr)
			// Don’t stop the walk, just log
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), strings.ToLower(fileSuffix)) {
			logger.Println("--------------------------------------------------")
			logger.Printf("Processing file: %s", info.Name())

			hashes, hashErr := calculateHashes(path)
			if hashErr != nil {
				logger.Printf("Error calculating hashes: %v", hashErr)
				return nil
			}
			md5Lower := strings.ToLower(hashes.MD5)
			sha1Lower := strings.ToLower(hashes.SHA1)

			logger.Printf("MD5:  %s", md5Lower)
			logger.Printf("SHA1: %s", sha1Lower)

			found := false
			var romInfo RomInfo
			if r, ok := md5Map[md5Lower]; ok {
				found = true
				romInfo = r
			} else if r, ok := sha1Map[sha1Lower]; ok {
				found = true
				romInfo = r
			}

			currentName := info.Name()
			if found {
				// remove leading numbers in known game ROM names
				desiredName := removeLeadingNumbers(romInfo.RomName)
				if currentName == desiredName {
					logger.Printf("File already named correctly: %s", currentName)
				} else {
					newPath := filepath.Join(filepath.Dir(path), desiredName)
					logger.Printf("Renaming '%s' to '%s'", currentName, desiredName)
					if err := os.Rename(path, newPath); err != nil {
						logger.Printf("Error renaming: %v", err)
					}
				}
			} else {
				// not found in .dat, append " - HACK" to the filename
				newName := addHRSuffix(currentName)
				if newName == currentName {
					logger.Printf("Not found in .dat and already contains '- HACK': %s", currentName)
				} else {
					newPath := filepath.Join(filepath.Dir(path), newName)
					logger.Printf("Not found in .dat. Renaming '%s' to '%s'", currentName, newName)
					if err := os.Rename(path, newPath); err != nil {
						logger.Printf("Error renaming: %v", err)
					}
				}
			}
			logger.Println("--------------------------------------------------")
		}
		return nil
	})

	if err != nil {
		logger.Printf("Error walking directory: %v\n", err)
		return err
	}

	logger.Println("Process completed.")
	return nil
}

func loadDatFile(datPath string) (*DataFile, error) {
	f, err := os.Open(datPath)
	if err != nil {
		return nil, fmt.Errorf("error opening .dat file: %v", err)
	}
	defer f.Close()

	var data DataFile
	if err := xml.NewDecoder(f).Decode(&data); err != nil {
		return nil, fmt.Errorf("error parsing .dat file: %v", err)
	}
	return &data, nil
}

func buildRomMaps(data *DataFile) (map[string]RomInfo, map[string]RomInfo) {
	md5Map := make(map[string]RomInfo)
	sha1Map := make(map[string]RomInfo)

	for _, game := range data.Games {
		for _, rom := range game.Roms {
			romMD5 := strings.ToLower(rom.MD5)
			romSHA1 := strings.ToLower(rom.SHA1)
			if romMD5 != "" {
				md5Map[romMD5] = RomInfo{GameName: game.Name, RomName: rom.Name}
			}
			if romSHA1 != "" {
				sha1Map[romSHA1] = RomInfo{GameName: game.Name, RomName: rom.Name}
			}
		}
	}
	return md5Map, sha1Map
}

func calculateHashes(filePath string) (FileHashes, error) {
	var result FileHashes

	file, err := os.Open(filePath)
	if err != nil {
		return result, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Calculate MD5
	md5Hasher := md5.New()
	if _, err := io.Copy(md5Hasher, file); err != nil {
		return result, fmt.Errorf("error calculating MD5: %v", err)
	}
	result.MD5 = hex.EncodeToString(md5Hasher.Sum(nil))

	// Reset file pointer for SHA-1
	if _, err := file.Seek(0, 0); err != nil {
		return result, fmt.Errorf("error resetting file pointer: %v", err)
	}
	sha1Hasher := sha1.New()
	if _, err := io.Copy(sha1Hasher, file); err != nil {
		return result, fmt.Errorf("error calculating SHA-1: %v", err)
	}
	result.SHA1 = hex.EncodeToString(sha1Hasher.Sum(nil))

	return result, nil
}

func removeLeadingNumbers(romName string) string {
	return reLeadingNumbers.ReplaceAllString(romName, "")
}

func addHRSuffix(currentName string) string {
	ext := filepath.Ext(currentName)
	base := strings.TrimSuffix(currentName, ext)
	return fmt.Sprintf("%s - HACK%s", base, ext)
}
