const file = Bun.file("input.txt");
const raw = await file.text();
const lines: string[] = raw.trim().split("\n");

type point = { x: number; y: number };
type claim = { id: number; x: number; y: number; w: number; h: number };
type grid = Record<string, Array<claim>>;

const g: grid = {};

const twoOrMoreOverlap = (g: grid): number =>
  Object.values(g).filter((claims) => claims.length >= 2).length;

const lookForIntactclaim = (g: grid): number => {
  for (const claims of Object.values(g)) {
    if (claims.length >= 2) continue;
    const claim = claims[0];
    let intact = true;
    for (let i = claim.x; i < claim.x + claim.w; i++) {
      for (let j = claim.y; j < claim.y + claim.h; j++) {
        const key = `${i},${j}`;
        if (g[key].length >= 2) {
          intact = false;
          break;
        }
      }
      if (!intact) break;
    }
    if (intact) return claim.id;
  }
  return -1;
};

lines.forEach((line) => {
  const [id, x, y, w, h] = line.match(/\d+/g)!.map(Number);
  const c: claim = { id, x, y, w, h };
  for (let i = x; i < x + w; i++) {
    for (let j = y; j < y + h; j++) {
      const key = `${i},${j}`;
      if (!g[key]) g[key] = [];
      g[key].push(c);
    }
  }
});

console.log(`Part 1: ${twoOrMoreOverlap(g)}`);
console.log(`Part 2: ${lookForIntactclaim(g)}`);
