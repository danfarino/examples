import { signal } from "@preact/signals";

export type Point = { x: number; y: number };

export const mousePos = signal<Point>({ x: 0, y: 0 });

function handler(e: MouseEvent) {
  mousePos.value = { x: e.clientX, y: e.clientY };
}

window.addEventListener("mousemove", handler, true);
