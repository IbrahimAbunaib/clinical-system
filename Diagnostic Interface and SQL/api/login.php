<?php
header('Content-Type: application/json');
require_once '../config/db_connect.php';

error_reporting(E_ALL);
ini_set('display_errors', 1);

try {
    $data = json_decode(file_get_contents('php://input'), true);
    $role = $data['role'] ?? '';
    $password = $data['password'] ?? '';

    if ($role === 'doctor') {
        $username = $data['username'] ?? '';
        if (empty($username)) {
            throw new Exception('Username is required');
        }
        $stmt = $conn->prepare("SELECT id, username, password FROM doctors WHERE username = ?");
        $stmt->bind_param("s", $username);
    } else {
        $patientId = $data['patientId'] ?? '';
        if (empty($patientId)) {
            throw new Exception('Patient ID is required');
        }
        $stmt = $conn->prepare("SELECT id, name as username, password FROM patients WHERE id = ?");
        $stmt->bind_param("s", $patientId);
    }

    if (!$stmt->execute()) {
        throw new Exception('Database query failed: ' . $stmt->error);
    }

    $result = $stmt->get_result();
    $user = $result->fetch_assoc();

    if (!$user) {
        throw new Exception('User not found');
    }

    if (!password_verify($password, $user['password'])) {
        // For doctor login with default credentials
        if ($role === 'doctor' && $username === 'doctor' && $password === 'doctor123') {
            echo json_encode([
                'success' => true,
                'user' => [
                    'id' => $user['id'],
                    'username' => $user['username']
                ]
            ]);
            exit;
        }
        throw new Exception('Invalid password');
    }

    // Remove password from response
    unset($user['password']);
    
    echo json_encode([
        'success' => true,
        'user' => $user
    ]);

} catch (Exception $e) {
    error_log("Login error: " . $e->getMessage());
    echo json_encode([
        'success' => false,
        'message' => $e->getMessage()
    ]);
}

if (isset($stmt)) {
    $stmt->close();
}
$conn->close();
?>
