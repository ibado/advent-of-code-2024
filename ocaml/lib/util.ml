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
