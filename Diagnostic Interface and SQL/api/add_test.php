<?php
header('Content-Type: application/json');
require_once '../includes/functions.php';

try {
    $data = json_decode(file_get_contents('php://input'), true);
    
    // Medical staff input validation
    if (empty($data['patientId']) || empty($data['testType']) || empty($data['result'])) {
        throw new Exception('Please fill in all required medical information');
    }

    // Verify patient exists in system
    if (!validatePatient($data['patientId'])) {
        throw new Exception('Patient not found in system. Please verify the ID.');
    }

    // Record the medical test
    $success = recordMedicalTest(
        $data['patientId'],
        $data['doctorId'],
        $data['testType'],
        $data['result'],
        $data['diagnosticsId'] ?? null,
        $data['recommendedAction'] ?? null
    );

    if ($success) {
        echo json_encode(['success' => true, 'message' => 'Medical test recorded successfully']);
    } else {
        throw new Exception('Failed to record medical test');
    }

} catch (Exception $e) {
    echo json_encode([
        'success' => false, 
        'message' => $e->getMessage()
    ]);
}
?>
