<?php
$host = "localhost";
$user = "root";
$pass = ""; // Change if your MySQL has a password
$dbname = "clinic_db";

$conn = new mysqli($host, $user, $pass, $dbname);

if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}
?>
