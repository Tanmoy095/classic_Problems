// 📝 Task : Implement a Logging Job Queue System

// Scenario:
// You are building a job processing system where different types of jobs are added to a queue and processed concurrently. Each job should log its execution details using a Logger interface.

// 🔹 Requirements

// Define a Job interface with a Execute() method.
// Create two job types (EmailJob and DataProcessingJob) that implement the Job interface.
// EmailJob: Simulates sending an email.
// DataProcessingJob: Simulates processing some data.
// Define a Logger interface with a Log(message string) method.
// Create a ConsoleLogger struct that implements the Logger interface.
// Create a JobProcessor struct that:
// Has a job queue ([]Job).
// Has a logger (Logger interface, passed via dependency injection).
// Uses goroutines and a WaitGroup to process jobs concurrently.
// In main(), do the following:
// Initialize ConsoleLogger.
// Add multiple jobs (EmailJob, DataProcessingJob) to JobProcessor.
// Process jobs concurrently and log their execution.
// 🖥️ Expected Output (Example)

// [LOG]: Processing Email Job: Sending email to user@example.com
// [LOG]: Processing Data Processing Job: Calculating user statistics...
// [LOG]: Email Job Completed
// [LOG]: Data Processing Job Completed

package main

import "fmt"

type Job interface {
	Execute()
}

//Email Job

type EmailJob struct {
}

func NewEmailJob() *EmailJob {
	return &EmailJob{}
}
func (job *EmailJob) Execute() {
	fmt.Println(" Processing Email Job")
}

//Data Processing Job

type DataProcessingJob struct {
}

func NewDataProcessingJob() *DataProcessingJob {
	return &DataProcessingJob{}
}

func (job *DataProcessingJob) Execute() {
	fmt.Println("Processing Data Processing Job")
}

//loggers.....

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type JobProcessor struct {
	logger Logger
	jobs   []Job
}

func NewJobProcessor(logger Logger) *JobProcessor {
	return &JobProcessor{logger: logger}
}
