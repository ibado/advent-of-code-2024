open Point

let count_trues (arr : bool array array) =
  Array.fold_left
    (Array.fold_left (fun acc it -> if it then acc + 1 else acc))
    0 arr

let find_point mx c =
  let i =
    Option.get @@ Array.find_index (fun s -> Array.exists (Char.equal c) s) mx
  in
  let j = Option.get @@ Array.find_index (Char.equal c) mx.(i) in
  Point.create i j

let parse_matrix input =
  Array.of_list input |> Array.map String.to_seq |> Array.map Array.of_seq

let index_of p arr =
  Option.get (Array.find_index (fun e -> e.x = p.x && e.y = p.y) arr)

let find_seen mx p dir =
  let rows, cols = (Array.length mx, Array.length mx.(0)) in
  let seen = Array.make_matrix rows cols false in
  let _ = seen.(p.x).(p.y) <- true in
  let rec aux mx p dir seen =
    let np = Point.add p dir in
    if not (Point.in_range rows cols np) then seen
    else
      match mx.(np.x).(np.y) with
      | '#' -> aux mx p (Point.rotate_right dir) seen
      | _ ->
          seen.(np.x).(np.y) <- true;
          aux mx np dir seen
  in
  aux mx p dir seen

let part1 input =
  let mx = parse_matrix input in
  let start = find_point mx '^' in
  count_trues @@ find_seen mx start { x = -1; y = 0 }

let part2 input =
  let mx = parse_matrix input in
  let rows, cols = (Array.length mx, Array.length mx.(0)) in
  let starting_dir = Point.create (-1) 0 in
  let is_loop mx p dir =
    let seen = Array.make (rows * cols * 4) false in
    let rec aux mx p dir =
      let np = Point.add p dir in
      if not (Point.in_range rows cols np) then false
      else
        match mx.(np.x).(np.y) with
        | '#' -> aux mx p (Point.rotate_right dir)
        | _ ->
            let idx =
              (((np.x * rows) + np.y) * 4) + index_of dir Point.ortogonals
            in
            if seen.(idx) then true
            else
              let () = seen.(idx) <- true in
              aux mx np dir
    in
    aux mx p dir
  in

  let start = find_point mx '^' in
  let seen_part1 = find_seen mx start starting_dir in

  let count = ref 0 in
  for i = 0 to rows - 1 do
    for j = 0 to cols - 1 do
      if seen_part1.(i).(j) then (
        mx.(i).(j) <- '#';
        if is_loop mx start starting_dir then incr count;
        mx.(i).(j) <- '.')
    done
  done;
  !count
