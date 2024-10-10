const IP = process.env['IPv4'];
const BASE_URL = `http://${IP}:8000/api/auth`;

export class CredentialsService {
    //? REGISTER
    static async register(email: string, password: string) {
        try {
            const response = await fetch(`${BASE_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "user": {
                        email,
                        password,
                    }
                }),
            });
            if (response.status !== 200) {
                console.log('register failed', response.status);
                throw new Error('Something went wrong.');
            }
            console.log('successful register');
            return;
        } catch (error) {
            console.error("Catched Register Error:", error);
            throw error;
        }
    }

    //? LOGIN
    static async login(email: string, password: string) {
        try {
            const response = await fetch(`${BASE_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    email,
                    password
                }),
            });

            if (response.status !== 200) {
                console.log('login failed', response.status);
                throw new Error('Incorrect username or password.');
            }
            console.log('successful login');
            return;
        } catch (error) {
            console.error("Catched Login Error:", error);
            throw error;
        }
    }
}
