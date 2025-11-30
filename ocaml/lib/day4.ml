let parse_input input =
  let mx =
    Array.of_list input |> Array.map String.to_seq |> Array.map Array.of_seq
  in
  let rows = Array.length mx in
  let cols = Array.length mx.(0) in
  (mx, rows, cols)

let part1 input =
  let mx, rows, cols = parse_input input in
  let count = ref 0 in
  let is_xmas p0 dir mx =
    let p1 = Point.add p0 dir in
    let p2 = Point.add p1 dir in
    let p3 = Point.add p2 dir in
    Point.in_range rows cols p3
    && mx.(p0.x).(p0.y) = 'X'
    && mx.(p1.x).(p1.y) = 'M'
    && mx.(p2.x).(p2.y) = 'A'
    && mx.(p3.x).(p3.y) = 'S'
  in
  for i = 0 to rows - 1 do
    for j = 0 to cols - 1 do
      for di = 0 to Array.length Point.all_dirs - 1 do
        let p = Point.create i j in
        let d = Point.all_dirs.(di) in
        if is_xmas p d mx then count := !count + 1
      done
    done
  done;
  !count

let part2 input =
  let mx, rows, cols = parse_input input in
  let is_x_mas p =
    let count = ref 0 in
    for i = 0 to 3 do
      let dir = Point.diagonals.(i) in
      let p0 = Point.minus p dir in
      let p1 = Point.add p dir in
      if
        Point.in_range rows cols p0
        && Point.in_range rows cols p1
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
