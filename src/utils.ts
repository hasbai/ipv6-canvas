import { Point } from "./models/point";

export function isValidCoordinate(current: Point, imageSize: Point): boolean {
  return current.ge(Point.zero()) && imageSize.ge(current);
}
