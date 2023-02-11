export class Point {
  constructor(public x: number, public y: number) {}

  toString(): string {
    return `(${this.x}, ${this.y})`;
  }

  static zero(): Point {
    return new Point(0, 0);
  }

  copy(): Point {
    return new Point(this.x, this.y);
  }

  set(point: Point) {
    this.x = point.x;
    this.y = point.y;
  }

  add(point: Point | number): Point {
    if (typeof point === "number") {
      this.x += point;
      this.y += point;
    } else {
      this.x += point.x;
      this.y += point.y;
    }
    return this;
  }

  sub(point: Point | number): Point {
    if (typeof point === "number") {
      this.x -= point;
      this.y -= point;
    } else {
      this.x -= point.x;
      this.y -= point.y;
    }
    return this;
  }

  times(point: Point | number): Point {
    if (typeof point === "number") {
      this.x *= point;
      this.y *= point;
    } else {
      this.x *= point.x;
      this.y *= point.y;
    }
    return this;
  }

  div(point: number): Point {
    this.x /= point;
    this.y /= point;
    return this;
  }

  floor(): Point {
    this.x = Math.floor(this.x);
    this.y = Math.floor(this.y);
    return this;
  }

  ceil(): Point {
    this.x = Math.ceil(this.x);
    this.y = Math.ceil(this.y);
    return this;
  }

  equals(point: Point): boolean {
    return this.x === point.x && this.y === point.y;
  }

  ge(point: Point): boolean {
    return this.x >= point.x && this.y >= point.y;
  }
}
