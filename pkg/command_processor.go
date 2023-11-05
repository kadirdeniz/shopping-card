package pkg

import (
	"encoding/json"
	"io/ioutil"
	"trendyol/internal/dto"
)

const (
	CommandsInputFileName  = "commands.json"
	ResponsesInputFileName = "responses.json"
)

type FileCommandProcessor struct {
	inputFilePath, outputFilePath string
	responses                     []dto.Response
}

type CommandProcessor interface {
	ReadFile() ([]dto.Request, error)
	WriteFile() error
	AppendResponse(dto.Response)
}

func NewFileCommandProcessor(inputFilePath, outputFilePath string) *FileCommandProcessor {
	return &FileCommandProcessor{
		inputFilePath:  inputFilePath,
		outputFilePath: outputFilePath,
	}
}

func (c *FileCommandProcessor) ReadFile() ([]dto.Request, error) {
	var data []dto.Request

	file, err := ioutil.ReadFile(c.inputFilePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *FileCommandProcessor) WriteFile() error {
	file, _ := json.MarshalIndent(c.responses, "", " ")
	err := ioutil.WriteFile(c.outputFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *FileCommandProcessor) AppendResponse(response dto.Response) {
	c.responses = append(c.responses, response)
}
