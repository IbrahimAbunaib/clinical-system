<?php
include 'db_connect.php';

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $name = $_POST['name'];
    $email = $_POST['email'];
    $phone = $_POST['phone'];
    $password = password_hash($_POST['password'], PASSWORD_DEFAULT); 

    $query = "INSERT INTO patients (name, email, phone, password) 
              VALUES ('$name', '$email', '$phone', '$password')";

    if ($conn->query($query) === TRUE) {
        echo "DONE";
    } else {
        echo "Error " . $conn->error;
    }
}
?>
