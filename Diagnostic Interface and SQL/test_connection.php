<?php
// Enable error reporting
error_reporting(E_ALL);
ini_set('display_errors', 1);
?>
<!DOCTYPE html>
<html>
<head>
    <title>Database Connection Test</title>
    <style>
        body { font-family: Arial, sans-serif; padding: 20px; }
        .success { color: green; }
        .error { color: red; }
    </style>
</head>
<body>
    <h2>Database Connection Test</h2>
    <?php
    try {
        require_once 'config/db_connect.php';

        if ($conn->connect_error) {
            throw new Exception("Connection failed: " . $conn->connect_error);
        }
        
        echo "<p class='success'>Database connection successful!</p>";
        
        // Test if tables exist
        $tables = ['doctors', 'patients', 'test_results'];
        foreach ($tables as $table) {
            $result = $conn->query("SHOW TABLES LIKE '$table'");
            if ($result->num_rows > 0) {
                echo "<p class='success'>Table '$table' exists</p>";
                
                // Show count of records
                $count = $conn->query("SELECT COUNT(*) as count FROM $table")->fetch_assoc()['count'];
                echo "<p>Number of records in $table: $count</p>";
            } else {
                echo "<p class='error'>Table '$table' does not exist</p>";
            }
        }
    } catch (Exception $e) {
        echo "<p class='error'>Error: " . $e->getMessage() . "</p>";
    } finally {
        if (isset($conn)) {
            $conn->close();
        }
    }
    ?>
</body>
</html>
