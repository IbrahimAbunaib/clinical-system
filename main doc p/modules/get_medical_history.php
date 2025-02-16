<?php
include 'db_connect.php';

$patient_id = 1; 

$query = "SELECT * FROM medical_history WHERE patient_id = '$patient_id'";
$result = $conn->query($query);

if ($result->num_rows > 0) {
    while ($row = $result->fetch_assoc()) {
        echo "<p>Diagnosis: " . $row['diagnosis'] . "</p>";
        echo "<p>Prescription: " . $row['prescription'] . "</p>";
        echo "<hr>";
    }
} else {
    echo "error";
}
?>
