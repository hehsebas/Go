// frontend/script.js
document.getElementById('loginForm').addEventListener('submit', function(e) {
    e.preventDefault(); // Evitar que el formulario se envíe de manera predeterminada

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errorMessage = document.getElementById('errorMessage');

    if (username === '' || password === '') {
        errorMessage.textContent = 'Por favor ingresa usuario y contraseña.';
    } else {
        // Aquí puedes hacer una llamada al backend para verificar las credenciales
        // Ejemplo de llamada usando fetch:
        fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                window.location.href = '/dashboard'; // Redirige al dashboard después del login
            } else {
                errorMessage.textContent = 'Credenciales incorrectas.';
            }
        })
        .catch(err => {
            console.error('Error al autenticar:', err);
            errorMessage.textContent = 'Hubo un error al intentar iniciar sesión.';
        });
    }
});
