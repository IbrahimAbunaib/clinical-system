<?php
include 'db_connect.php';

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $patient_id = $_POST['patient_id'];
    $doctor_id = $_POST['doctor_id'];
    $appointment_date = $_POST['appointment_date'];

    $query = "INSERT INTO appointments (patient_id, doctor_id, appointment_date) 
              VALUES ('$patient_id', '$doctor_id', '$appointment_date')";

    if ($conn->query($query) === TRUE) {
        echo "DONE";
    } else {
        echo "Error " . $conn->error;
    }
}
?>
