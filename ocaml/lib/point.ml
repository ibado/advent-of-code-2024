type point = { x : int; y : int }

let create x y = { x; y }

let ortogonals : point array =
  [| { x = -1; y = 0 }; { x = 1; y = 0 }; { x = 0; y = -1 }; { x = 0; y = 1 } |]

let diagonals : point array =
  [|
    { x = -1; y = -1 }; { x = -1; y = 1 }; { x = 1; y = -1 }; { x = 1; y = 1 };
  |]

let all_dirs = Array.append ortogonals diagonals
let add p1 p2 = { x = p1.x + p2.x; y = p1.y + p2.y }
let in_range xmax ymax p = p.x >= 0 && p.x < xmax && p.y >= 0 && p.y < ymax
let minus p1 p2 = { x = p1.x - p2.x; y = p1.y - p2.y }
let rotate_right p = { x = p.y; y = -p.x }
