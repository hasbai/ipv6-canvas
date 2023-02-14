import { Point } from "@/models/point";

export class Color {
  constructor(
    public r: number,
    public g: number,
    public b: number,
    public a: number
  ) {}

  fromCanvas(ctx: CanvasRenderingContext2D, coordinate: Point) {
    const imageData = ctx.getImageData(coordinate.x, coordinate.y, 1, 1);
    this.r = imageData.data[0];
    this.g = imageData.data[1];
    this.b = imageData.data[2];
    this.a = imageData.data[3];
  }

  toString(): string {
    return (
      "#" +
      this.r.toString(16).padStart(2, "0") +
      this.g.toString(16).padStart(2, "0") +
      this.b.toString(16).padStart(2, "0") +
      this.a.toString(16).padStart(2, "0")
    );
  }
}

export class Pixel {
    constructor(public coordinate: Point, public color: Color) {
    }

    static fromBytes(data: Uint8Array): Pixel {
        return new Pixel(
            new Point((data[1] << 8) | data[2], (data[3] << 8) | data[4]),
            new Color(data[5], data[6], data[7], data[8])
        );
    }
}
