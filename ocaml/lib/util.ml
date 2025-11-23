let read_lines filename =
  In_channel.with_open_text filename In_channel.input_lines

let input_day day =
  assert (day < 25);
  read_lines (Printf.sprintf "../input/%d.txt" day)

let parse_nums line =
  String.split_on_char ' ' line
  |> List.filter_map (fun s ->
         if s <> "" then Option.bind (Some s) int_of_string_opt else None)

let rec window2 f acc l =
  match l with
  | [] -> acc
  | [ _ ] -> acc
  | n1 :: n2 :: l -> (
      match f acc n1 n2 with None -> acc | Some r -> window2 f r ([ n2 ] @ l))

(* Point utils *)

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
let point_rotate_right p = { x = p.y; y = -p.x }
