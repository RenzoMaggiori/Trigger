const IP = process.env['IPv4'];
const BASE_URL = `http://${IP}:8080/api/auth`;

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
                throw new Error('Register failed');
            }
            return response.json();
            
        } catch (error) {
            console.error("Register error:", error);
            throw error;
        }
    }
}