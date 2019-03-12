export class Speaker {
    id: number;
    firstName: string;
    lastName: string;
    email: string;

    constructor(firstName: string, lastName: string, email: string) {
        this.id = -1;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
}
}