<?php
include 'db_connect.php';

$doctor_id = 1; 
$type = $_GET['type'];

if ($type == "schedule") {
    $query = "SELECT schedule FROM doctor_schedule WHERE doctor_id = '$doctor_id'";
} elseif ($type == "consultation") {
    $query = "SELECT consultation FROM consultations WHERE doctor_id = '$doctor_id'";
} elseif ($type == "diagnostics") {
    $query = "SELECT test_request, test_result FROM diagnostics WHERE doctor_id = '$doctor_id'";
} elseif ($type == "prescription") {
    $query = "SELECT prescription FROM prescriptions WHERE doctor_id = '$doctor_id'";
}

$result = $conn->query($query);

if ($result->num_rows > 0) {
    while ($row = $result->fetch_assoc()) {
        echo "<p>" . implode(" - ", $row) . "</p>";
    }
} else {
    echo "Done";
}
?>
