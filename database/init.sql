-- Create the database (if it does not already exist)
CREATE DATABASE IF NOT EXISTS PetRescue;

-- Use the PetRescue database
USE PetRescue;

-- Create a table for users
CREATE TABLE Users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(255) NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create a table for pets
CREATE TABLE Pets (
    PetID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Type VARCHAR(50),
    Description TEXT,
    Status VARCHAR(50) DEFAULT 'available', -- e.g., available, adopted, in treatment
    ImageURL VARCHAR(255),
    ReportedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create a table for services offered by the platform
CREATE TABLE Services (
    ServiceID INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL
);

-- Create a table for volunteer opportunities
CREATE TABLE VolunteerOpportunities (
    OpportunityID INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL,
    AvailablePositions INT DEFAULT 1,
    Active BOOLEAN DEFAULT TRUE
);

-- Create a table for donations
CREATE TABLE Donations (
    DonationID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT,
    Amount DECIMAL(10, 2) NOT NULL,
    DonatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);
