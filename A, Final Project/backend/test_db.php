<?php
$host = "localhost";
$username = "root";
$password = "";
$database = "clinic_system";

$conn = new mysqli($host, $username, $password, $database);


if ($conn->connect_error) {
    die("Error " . $conn->connect_error);
} else {
    echo "DONE";
}

$query = "SELECT * FROM doctors";
$result = $conn->query($query);

if ($result->num_rows > 0) {
    while ($row = $result->fetch_assoc()) {
        echo "<p>Doctor: " . $row['name'] . " - " . $row['email'] . "</p>";
    }
} else {
    echo "Error";
}

$conn->close();
?>
