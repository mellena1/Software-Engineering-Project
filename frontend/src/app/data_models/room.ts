export class Room {
  id: number;
  capacity: number;
  name: string;

  constructor(name: string, capacity: number) {
    this.id = -1;
    this.capacity = capacity;
    this.name = name;
  }
}
