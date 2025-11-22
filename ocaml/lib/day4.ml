type point = { x : int; y : int }

let ortogonals : point array =
  [| { x = -1; y = 0 }; { x = 1; y = 0 }; { x = 0; y = -1 }; { x = 0; y = 1 } |]

let diagonals : point array =
  [|
    { x = -1; y = -1 }; { x = -1; y = 1 }; { x = 1; y = -1 }; { x = 1; y = 1 };
  |]

let all_dirs = Array.append ortogonals diagonals
let point_add p1 p2 = { x = p1.x + p2.x; y = p1.y + p2.y }
let point_in_range r p = p.x >= 0 && p.x < r.x && p.y >= 0 && p.y < r.y
let point_minus p1 p2 = { x = p1.x - p2.x; y = p1.y - p2.y }

let parse_input input =
  let mx =
    Array.of_list input |> Array.map String.to_seq |> Array.map Array.of_seq
  in
  let rows = Array.length mx in
  let cols = Array.length mx.(0) in
  (mx, rows, cols)

let part1 input =
  let mx, rows, cols = parse_input input in
  let range_point = { x = rows; y = cols } in
  let count = ref 0 in
  let is_xmas p0 dir mx =
    let p1 = point_add p0 dir in
    let p2 = point_add p1 dir in
    let p3 = point_add p2 dir in
    point_in_range range_point p3
    && mx.(p0.x).(p0.y) = 'X'
    && mx.(p1.x).(p1.y) = 'M'
    && mx.(p2.x).(p2.y) = 'A'
    && mx.(p3.x).(p3.y) = 'S'
  in
  for i = 0 to rows - 1 do
    for j = 0 to cols - 1 do
      for di = 0 to Array.length all_dirs - 1 do
        let p = { x = i; y = j } in
        let d = all_dirs.(di) in
        if is_xmas p d mx then count := !count + 1
      done
    done
  done;
  !count

let part2 input =
  let mx, rows, cols = parse_input input in
  let range_point = { x = rows; y = cols } in
  let is_x_mas p =
    let count = ref 0 in
    for i = 0 to 3 do
      let dir = diagonals.(i) in
      let p0 = point_minus p dir in
      let p1 = point_add p dir in
      if
        point_in_range range_point p0
        && point_in_range range_point p1
        && mx.(p0.x).(p0.y) = 'M'
        && mx.(p1.x).(p1.y) = 'S'
      then incr count
    done;
    !count = 2
  in
  let count = ref 0 in
  for i = 0 to rows - 1 do
    for j = 0 to cols - 1 do
      if mx.(i).(j) = 'A' && is_x_mas { x = i; y = j } then incr count
    done
  done;
  !count
