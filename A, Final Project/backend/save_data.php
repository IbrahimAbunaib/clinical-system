<?php
include 'db_connect.php';

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $doctor_id = 1; 

    if (isset($_POST['schedule'])) {
        $schedule = $_POST['schedule'];
        $query = "INSERT INTO doctor_schedule (doctor_id, schedule) VALUES ('$doctor_id', '$schedule')";
    } elseif (isset($_POST['consultation'])) {
        $consultation = $_POST['consultation'];
        $query = "INSERT INTO consultations (doctor_id, consultation) VALUES ('$doctor_id', '$consultation')";
    } elseif (isset($_POST['test_request'])) {
        $test_request = $_POST['test_request'];
        $test_result = $_POST['test_result'];
        $query = "INSERT INTO diagnostics (doctor_id, test_request, test_result) VALUES ('$doctor_id', '$test_request', '$test_result')";
    } elseif (isset($_POST['prescription'])) {
        $prescription = $_POST['prescription'];
        $query = "INSERT INTO prescriptions (doctor_id, prescription) VALUES ('$doctor_id', '$prescription')";
    }

    if ($conn->query($query) === TRUE) {
        echo "Done";
    } else {
        echo "error " . $conn->error;
    }
}
?>