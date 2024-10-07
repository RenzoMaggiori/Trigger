const IP = process.env['IPv4'];
const BASE_URL = `http://${IP}:8080/api/auth`;

export class CredentialsService {
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
                throw new Error('Login failed');
            }
            return;
        } catch (error) {
            console.error("Catched Login Error:", error);
            throw error;
        }
    }
}