<?php
header('Content-Type: application/json');
require_once '../config/db_connect.php';

try {
    $data = json_decode(file_get_contents('php://input'), true);
    
    // Validate required fields
    if (empty($data['patientId']) || empty($data['testType']) || empty($data['result'])) {
        throw new Exception('Required fields missing');
    }

    // First check if patient exists
    $checkPatient = $conn->prepare("SELECT id FROM patients WHERE id = ?");
    $checkPatient->bind_param("s", $data['patientId']);
    $checkPatient->execute();
    $patientResult = $checkPatient->get_result();
    
    if ($patientResult->num_rows === 0) {
        throw new Exception('Patient ID does not exist. Please verify the Patient ID.');
    }
    $checkPatient->close();

    // Check if doctor exists
    $checkDoctor = $conn->prepare("SELECT id FROM doctors WHERE id = ?");
    $checkDoctor->bind_param("s", $data['doctorId']);
    $checkDoctor->execute();
    $doctorResult = $checkDoctor->get_result();
    
    if ($doctorResult->num_rows === 0) {
        throw new Exception('Invalid doctor ID');
    }
    $checkDoctor->close();

    // Insert test result
    $stmt = $conn->prepare("INSERT INTO test_results (patient_id, doctor_id, test_type, diagnostics_id, result, recommended_action) VALUES (?, ?, ?, ?, ?, ?)");
    
    $stmt->bind_param("ssssss", 
        $data['patientId'],
        $data['doctorId'],
        $data['testType'],
        $data['diagnosticsId'],
        $data['result'],
        $data['recommendedAction']
    );
    
    if ($stmt->execute()) {
        echo json_encode(['success' => true, 'message' => 'Test result added successfully']);
    } else {
        throw new Exception('Failed to add test result: ' . $stmt->error);
    }
    
    $stmt->close();

} catch (Exception $e) {
    echo json_encode([
        'success' => false, 
        'message' => $e->getMessage(),
        'details' => 'Make sure the Patient ID exists in the system'
    ]);
}

$conn->close();
?>
