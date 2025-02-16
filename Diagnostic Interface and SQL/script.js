let currentUser = null;

// Basic UI functions
function showTab(tabName) {
    document.querySelectorAll('.tab-content').forEach(tab => tab.style.display = 'none');
    document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
    document.getElementById(tabName + 'Tab').style.display = 'block';
    document.querySelector(`[onclick="showTab('${tabName}')"]`).classList.add('active');
}

function toggleLoginFields() {
    const role = document.getElementById('role').value;
    document.getElementById('doctorFields').style.display = role === 'doctor' ? 'block' : 'none';
    document.getElementById('patientFields').style.display = role === 'patient' ? 'block' : 'none';
}

// Medical System Functions
async function addTestResult() {
    const testData = {
        patientId: document.getElementById('patientId').value,
        testType: document.getElementById('testType').value,
        result: document.getElementById('testResult').value,
        diagnosticsId: document.getElementById('diagnosticsId').value,
        recommendedAction: document.getElementById('recommendedAction').value,
        doctorId: currentUser.id
    };

    if (!testData.patientId || !testData.testType || !testData.result) {
        alert('Please fill in all required medical information');
        return;
    }

    try {
        const response = await fetch('api/add_test.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(testData)
        });

        const result = await response.json();
        if (result.success) {
            alert('Medical test recorded successfully');
            document.querySelectorAll('#testForm input, #testForm textarea').forEach(input => input.value = '');
            updateTestList();
        } else {
            alert(result.message);
        }
    } catch (error) {
        alert('System Error: Unable to record medical test');
        console.error(error);
    }
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

// Registration
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

async function updateTestList() {
    if (!currentUser) return;

    try {
        const response = await fetch('api/get_tests.php', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                role: currentUser.role,
                userId: currentUser.id
            })
        });

        const data = await response.json();
        const testList = document.getElementById('testList');
        
        testList.innerHTML = data.success && data.tests.length > 0 
            ? data.tests.map(test => `
                <div class="test-item">
                    <h4>${test.test_type.toUpperCase()} - Patient: ${test.patient_name}</h4>
                    <p>Result: ${test.result}</p>
                    <p>Diagnostics ID: ${test.diagnostics_id || 'N/A'}</p>
                    <p>Recommended Action: ${test.recommended_action || 'N/A'}</p>
                    <p>Date: ${new Date(test.created_at).toLocaleString()}</p>
                    <p>Doctor: ${test.doctor_name}</p>
                </div>
            `).join('')
            : '<p>No medical records found</p>';
    } catch (error) {
        document.getElementById('testList').innerHTML = '<p>Error loading medical records</p>';
        console.error(error);
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