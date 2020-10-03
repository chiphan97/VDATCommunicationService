import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class GenerateColorService {

  private colors: Set<string>;

  constructor() {
    this.colors = new Set<string>();
  }

  public generate(): string {
    let color = '#';
    const characters = '0123456789ABCDEF';
    const charactersLength = characters.length;

    for (let i = 0; i < 6; i++) {
      color += characters.charAt(Math.floor(Math.random() * charactersLength));
    }

    if (this.colors.has(color)) {
      color = this.generate();
    }

    this.colors.add(color);
    return color;
  }
}
