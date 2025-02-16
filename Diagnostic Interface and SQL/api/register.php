<?php
header('Content-Type: application/json');
require_once '../config/db_connect.php';

try {
    $data = json_decode(file_get_contents('php://input'), true);
    
    // Validate required fields
    if (empty($data['patientId']) || empty($data['name']) || empty($data['password'])) {
        throw new Exception('All fields are required');
    }

    // Check if patient ID already exists
    $checkStmt = $conn->prepare("SELECT id FROM patients WHERE id = ?");
    $checkStmt->bind_param("s", $data['patientId']);
    $checkStmt->execute();
    if ($checkStmt->get_result()->num_rows > 0) {
        throw new Exception('Patient ID already exists');
    }
    $checkStmt->close();

    // Hash the password
    $hashedPassword = password_hash($data['password'], PASSWORD_DEFAULT);

    // Insert new patient
    $stmt = $conn->prepare("INSERT INTO patients (id, name, password) VALUES (?, ?, ?)");
    $stmt->bind_param("sss", $data['patientId'], $data['name'], $hashedPassword);
    
    if ($stmt->execute()) {
        echo json_encode(['success' => true, 'message' => 'Registration successful']);
    } else {
        throw new Exception('Registration failed');
    }
    $stmt->close();

} catch (Exception $e) {
    echo json_encode(['success' => false, 'message' => $e->getMessage()]);
}

$conn->close();
?>
