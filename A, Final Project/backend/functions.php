<?php
require_once '../config/db_connect.php';

function validatePatient($patientId) {
    global $conn;
    $stmt = $conn->prepare("SELECT id FROM patients WHERE id = ?");
    $stmt->bind_param("s", $patientId);
    $stmt->execute();
    $exists = $stmt->get_result()->num_rows > 0;
    $stmt->close();
    return $exists;
}

function recordMedicalTest($patientId, $doctorId, $testType, $result, $diagnosticsId = null, $recommendedAction = null) {
    global $conn;
    $stmt = $conn->prepare("INSERT INTO test_results (patient_id, doctor_id, test_type, result, diagnostics_id, recommended_action) VALUES (?, ?, ?, ?, ?, ?)");
    $stmt->bind_param("ssssss", $patientId, $doctorId, $testType, $result, $diagnosticsId, $recommendedAction);
    $success = $stmt->execute();
    $stmt->close();
    return $success;
}

function getMedicalHistory($patientId = null) {
    global $conn;
    $query = "SELECT 
                t.*, 
                p.name as patient_name, 
                d.username as doctor_name 
              FROM test_results t 
              JOIN patients p ON t.patient_id = p.id 
              JOIN doctors d ON t.doctor_id = d.id";
    
    if ($patientId) {
        $query .= " WHERE t.patient_id = ?";
        $stmt = $conn->prepare($query);
        $stmt->bind_param("s", $patientId);
    } else {
        $stmt = $conn->prepare($query);
    }
    
    $stmt->execute();
    $results = $stmt->get_result()->fetch_all(MYSQLI_ASSOC);
    $stmt->close();
    return $results;
}
?>
