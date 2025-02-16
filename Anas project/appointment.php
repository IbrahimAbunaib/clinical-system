<?php
$servername = "localhost";
$username = "root"; 
$password = ""; 
$dbname = "appointment";

$conn = new mysqli($servername, $username, $password, $dbname);

if ($conn->connect_error) {
    die("error:" . $conn->connect_error);
}
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $patient_id = $_POST['patient_id'];
    $doctor_id = $_POST['doctor_id'];
    $appointment_date = $_POST['appointment_date'];
    $appointment_time = $_POST['appointment_time'];
    $appointment_status = $_POST['appointment_status'];
    $created_at = $_POST['created_at'];

    $sql = "INSERT INTO appointments (patient_id, doctor_id, appointment_date, appointment_time, appointment_status, created_at) 
            VALUES ('$patient_id', '$doctor_id', '$appointment_date', '$appointment_time', '$appointment_status', '$created_at')";

    if ($conn->query($sql) === TRUE) {
        echo "<script>alert('successful'); window.location.href='index.html';</script>";
    } else {
        echo "faied : " . $conn->error;
    }
}

$conn->close();
?>