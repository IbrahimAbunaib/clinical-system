<?php
header('Content-Type: application/json');
require_once '../config/db_connect.php';

error_reporting(E_ALL);
ini_set('display_errors', 1);

try {
    $data = json_decode(file_get_contents('php://input'), true);
    $role = $data['role'] ?? '';
    $userId = $data['userId'] ?? '';

    // Debug log
    error_log("Fetching tests for role: " . $role . ", userId: " . $userId);

    $query = "SELECT t.*, p.name as patient_name, d.username as doctor_name 
              FROM test_results t 
              LEFT JOIN patients p ON t.patient_id = p.id 
              LEFT JOIN doctors d ON t.doctor_id = d.id";

    if ($role === 'patient') {
        $query .= " WHERE t.patient_id = ?";
        $stmt = $conn->prepare($query);
        $stmt->bind_param("s", $userId);
    } else {
        // For doctors, show all tests
        $stmt = $conn->prepare($query);
    }

    if (!$stmt->execute()) {
        throw new Exception("Query execution failed: " . $stmt->error);
    }

    $result = $stmt->get_result();
    $tests = [];
    
    while ($row = $result->fetch_assoc()) {
        $tests[] = $row;
    }

    // Debug log
    error_log("Found " . count($tests) . " test results");
    error_log("Tests data: " . json_encode($tests));

    echo json_encode(['success' => true, 'tests' => $tests]);

} catch (Exception $e) {
    error_log("Error in get_tests.php: " . $e->getMessage());
    echo json_encode(['success' => false, 'message' => $e->getMessage()]);
}

$conn->close();
?>
