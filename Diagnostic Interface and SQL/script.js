// Simulated database
let testResults = [];
let currentUser = null;
let registeredPatients = [];

function showTab(tabName) {
  document.querySelectorAll('.tab-content').forEach(tab => tab.style.display = 'none');
  document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));

  document.getElementById(tabName + 'Tab').style.display = 'block';
  document.querySelector(`[onclick="showTab('${tabName}')"]`).classList.add('active');
}

async function register() {
    const patientId = document.getElementById('regPatientId').value;
    const name = document.getElementById('regName').value;
    const password = document.getElementById('regPassword').value;
    const confirmPassword = document.getElementById('regConfirmPassword').value;

    if (!patientId || !name || !password) {
        alert('Please fill all fields');
        return;
    }

    if (password !== confirmPassword) {
        alert('Passwords do not match');
        return;
    }

    try {
        const response = await fetch('api/register.php', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                patientId,
                name,
                password
            })
        });

        const data = await response.json();
        
        if (data.success) {
            alert('Registration successful! Please login.');
            showTab('login');
            // Clear registration form
            document.getElementById('regPatientId').value = '';
            document.getElementById('regName').value = '';
            document.getElementById('regPassword').value = '';
            document.getElementById('regConfirmPassword').value = '';
        } else {
            alert(data.message || 'Registration failed');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred during registration');
    }
}

// Toggle login fields based on role
function toggleLoginFields() {
  const role = document.getElementById('role').value;
  document.getElementById('doctorFields').style.display = role === 'doctor' ? 'block' : 'none';
  document.getElementById('patientFields').style.display = role === 'patient' ? 'block' : 'none';
}

// User authentication
async function login() {
    const role = document.getElementById('role').value;
    const password = document.getElementById('loginPassword').value;
    
    let loginData = {
        role: role,
        password: password
    };

    if (role === 'doctor') {
        loginData.username = document.getElementById('doctorUsername').value;
    } else {
        loginData.patientId = document.getElementById('loginPatientId').value;
    }

    try {
        console.log('Sending login data:', loginData); // Debug log
        const response = await fetch('api/login.php', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(loginData)
        });

        const data = await response.json();
        console.log('Login response:', data); // Debug log
        
        if (data.success) {
            currentUser = { ...data.user, role };
            document.getElementById('loginSection').style.display = 'none';
            document.getElementById('mainContent').style.display = 'block';
            document.getElementById('userRole').textContent = `${role.toUpperCase()}: ${currentUser.username}`;
            document.getElementById('doctorView').style.display = role === 'doctor' ? 'block' : 'none';
            document.getElementById('patientView').style.display = role === 'patient' ? 'block' : 'none';
            updateTestList();
        } else {
            alert(data.message || 'Login failed');
        }
    } catch (error) {
        console.error('Login error:', error);
        alert('An error occurred during login. Please check the console for details.');
    }
}

// Add new test result
async function addTestResult() {
    const patientId = document.getElementById('patientId').value;
    const testType = document.getElementById('testType').value;
    const testResult = document.getElementById('testResult').value;
    const diagnosticsId = document.getElementById('diagnosticsId').value;
    const recommendedAction = document.getElementById('recommendedAction').value;

    if (!patientId || !testType || !testResult) {
        alert('Please fill in all required fields');
        return;
    }

    try {
        const response = await fetch('api/add_test.php', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                patientId,
                doctorId: currentUser.id,
                testType,
                result: testResult,
                diagnosticsId,
                recommendedAction
            })
        });

        const data = await response.json();
        
        if (data.success) {
            alert('Test result added successfully');
            clearTestForm();
            updateTestList();
        } else {
            alert(data.message || 'Failed to add test result');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred while adding the test result');
    }
}

// Update test list display
async function updateTestList() {
    const testList = document.getElementById('testList');
    testList.innerHTML = '<p>Loading test results...</p>';

    if (!currentUser) {
        console.log('No current user found');
        return;
    }

    try {
        const requestData = {
            role: currentUser.role,
            userId: currentUser.id
        };
        console.log('Fetching test results with data:', requestData);

        const response = await fetch('api/get_tests.php', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData)
        });

        const data = await response.json();
        console.log('Received test results:', data);
        
        if (data.success) {
            testList.innerHTML = '';
            
            if (!data.tests || data.tests.length === 0) {
                testList.innerHTML = '<p>No test results found.</p>';
                return;
            }

            data.tests.forEach(test => {
                const testElement = document.createElement('div');
                testElement.className = 'test-item';
                testElement.innerHTML = `
                    <h4>${test.test_type.toUpperCase()} - Patient: ${test.patient_name || test.patient_id}</h4>
                    <p>Result: ${test.result}</p>
                    <p>Diagnostics ID: ${test.diagnostics_id || 'N/A'}</p>
                    <p>Recommended Action: ${test.recommended_action || 'N/A'}</p>
                    <p>Date: ${new Date(test.created_at).toLocaleString()}</p>
                    <p>Doctor: ${test.doctor_name}</p>
                `;
                testList.appendChild(testElement);
            });
        } else {
            console.error('Error from server:', data.message);
            testList.innerHTML = '<p>Error loading test results: ' + (data.message || 'Unknown error') + '</p>';
        }
    } catch (error) {
        console.error('Error fetching test results:', error);
        testList.innerHTML = '<p>Error loading test results. Please try again later.</p>';
    }
}

// Clear test input form
function clearTestForm() {
  document.getElementById('patientId').value = '';
  document.getElementById('testResult').value = '';
  document.getElementById('diagnosticsId').value = '';
  document.getElementById('recommendedAction').value = '';
}

// Logout functionality
document.getElementById('logoutBtn').addEventListener('click', () => {
  currentUser = null;
  document.getElementById('loginSection').style.display = 'block';
  document.getElementById('mainContent').style.display = 'none';
  document.getElementById('username').value = '';
  document.getElementById('password').value = '';
  document.getElementById('testList').innerHTML = ''; // Clear test list on logout
  document.getElementById('reportOutput').innerHTML = ''; // Clear reports on logout
});