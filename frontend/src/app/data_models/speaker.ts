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

  static getFullName(speaker: Speaker): string {
    var name = "";
    var needSpace = false;
    if ("firstName" in speaker && speaker.firstName !== null) {
      name += speaker.firstName;
      needSpace = true;
    }
    if ("lastName" in speaker && speaker.lastName !== null) {
      if (needSpace) {
        name += " ";
      }
      name += speaker.lastName;
    }
    return name;
  }
}
