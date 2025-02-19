<?php
header('Content-Type: application/json');
require_once '../includes/functions.php';

error_reporting(E_ALL);
ini_set('display_errors', 1);

try {
    $data = json_decode(file_get_contents('php://input'), true);
    $patientId = ($data['role'] === 'patient') ? $data['userId'] : null;
    
    // Get medical history based on user role
    $medicalHistory = getMedicalHistory($patientId);
    
    echo json_encode([
        'success' => true, 
        'tests' => $medicalHistory
    ]);

} catch (Exception $e) {
    echo json_encode([
        'success' => false, 
        'message' => 'Unable to retrieve medical history: ' . $e->getMessage()
    ]);
}

$conn->close();
?>
