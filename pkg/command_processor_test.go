package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"trendyol/internal/dto"

	"github.com/google/go-cmp/cmp"
)

const (
	TestInputFile  = "test_commands_input.json"
	TestOutputFile = "test_commands_output.json"
)

// setupTestData prepares the test data and writes it to a file.
func setupTestData() []dto.Request {
	testData := []dto.Request{
		{
			Command: "addItem",
			Payload: dto.Payload{ItemID: 1, CategoryID: 3003, SellerID: 1, Quantity: 1, Price: 100},
		},
		{
			Command: "addVasItem",
			Payload: dto.Payload{ItemID: 1, CategoryID: 3003, SellerID: 1, Quantity: 1, Price: 100, VasItemId: 1},
		},
		{
			Command: "display",
			Payload: dto.Payload{},
		},
		{
			Command: "removeItem",
			Payload: dto.Payload{ItemID: 1},
		},
		{
			Command: "reset",
			Payload: dto.Payload{},
		},
	}
	file, _ := json.MarshalIndent(testData, "", " ")
	_ = ioutil.WriteFile(TestInputFile, file, 0644)
	return testData
}

// cleanupTestData removes the test files.
func cleanupTestData() {
	os.Remove(TestInputFile)
	os.Remove(TestOutputFile)
}

func TestReadFile(t *testing.T) {
	testData := setupTestData()
	defer cleanupTestData()

	// Make sure to read from the same file you wrote to in setupTestData
	processor := NewFileCommandProcessor(TestInputFile, TestOutputFile)
	requests, err := processor.ReadFile()

	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if diff := cmp.Diff(testData, requests); diff != "" {
		t.Errorf("ReadFile mismatch (-want +got):\n%s", diff)
	}
}

func TestReadFile_WrongJSONFormat(t *testing.T) {
	testData := []byte(`{"command": "addItem", "payload": {"item_id": 1, "category_id": 3003, "seller_id": 1, "quantity": 1, "price": 100}}`)
	_ = ioutil.WriteFile(TestInputFile, testData, 0644)
	defer cleanupTestData()

	processor := NewFileCommandProcessor(TestInputFile, TestOutputFile)
	_, err := processor.ReadFile()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestReadFile_NotFound(t *testing.T) {
	processor := NewFileCommandProcessor("nonexistent_file", TestOutputFile)
	_, err := processor.ReadFile()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestWriteFile(t *testing.T) {
	// Cleanup if a test file already exists
	cleanupTestData()

	testResponse := dto.Response{
		Result:  true,
		Message: "Item added successfully",
	}

	// Initialize a new FileCommandProcessor
	processor := NewFileCommandProcessor(TestInputFile, TestOutputFile)

	// Append the testResponse
	processor.AppendResponse(testResponse)

	// Write the response to file
	err := processor.WriteFile()
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// Validate the written data
	data, err := ioutil.ReadFile(TestOutputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var responses []dto.Response
	err = json.Unmarshal(data, &responses)
	if err != nil {
		t.Fatalf("Failed to unmarshal json data: %v", err)
	}

	if len(responses) != 1 || !responses[0].Result || responses[0].Message != "Item added successfully" {
		t.Errorf("Unexpected response written to file: %+v", responses)
	}

	// Cleanup after test
	cleanupTestData()
}

func TestWriteFile_Error(t *testing.T) {
	// Test with invalid filename
	processor := NewFileCommandProcessor(TestInputFile, "/invalid/path/test.json")

	err := processor.WriteFile()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAppendResponse(t *testing.T) {
	// Cleanup if a test file already exists
	cleanupTestData()

	testResponse := dto.Response{
		Result:  true,
		Message: "Item added successfully",
	}

	// Initialize a new FileCommandProcessor
	processor := NewFileCommandProcessor(TestInputFile, TestOutputFile)

	// Append the testResponse
	processor.AppendResponse(testResponse)

	// Explicitly write the file
	err := processor.WriteFile()
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// Validate the written data
	data, err := ioutil.ReadFile(TestOutputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var responses []dto.Response
	err = json.Unmarshal(data, &responses)
	if err != nil {
		t.Fatalf("Failed to unmarshal json data: %v", err)
	}

	if len(responses) != 1 || !responses[0].Result || responses[0].Message != "Item added successfully" {
		t.Errorf("Unexpected response written to file: %+v", responses)
	}

	// Cleanup after test
	cleanupTestData()
}
