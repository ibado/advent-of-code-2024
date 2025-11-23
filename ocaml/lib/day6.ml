open Util

let count_trues (arr : bool array array) =
  Array.fold_left
    (Array.fold_left (fun acc it -> if it then acc + 1 else acc))
    0 arr

let part1 input =
  let mx =
    Array.of_list input |> Array.map String.to_seq |> Array.map Array.of_seq
  in
  let rows = Array.length mx in
  let cols = Array.length mx.(0) in
  let rp = { x = rows; y = cols } in
  let i =
    Option.get @@ List.find_index (fun s -> String.contains s '^') input
  in
  let j = Option.get @@ Array.find_index (fun c -> c = '^') mx.(i) in
  let seen = Array.make_matrix rows cols false in
  let start : point = { x = i; y = j } in
  seen.(start.x).(start.y) <- true;
  let rec solve mx p seen dir =
    let np = point_add p dir in
    if not (point_in_range rp np) then count_trues seen
    else
      match mx.(np.x).(np.y) with
      | '#' -> solve mx p seen (point_rotate_right dir)
      | '.' ->
          seen.(np.x).(np.y) <- true;
          solve mx np seen dir
      | '^' -> solve mx np seen dir
      | _ -> raise (Invalid_argument "unreachable")
  in
  solve mx start seen { x = -1; y = 0 }

let part2 _input = 0
