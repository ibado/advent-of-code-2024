let product groups =
  int_of_string (Re.Group.get groups 1) * int_of_string (Re.Group.get groups 2)

let part1 input =
  let regexp = Re.Pcre.re {|mul\((\d+),(\d+)\)|} |> Re.compile in
  let rec muls ?(pos = 0) ?(acc = 0) s =
    match Re.exec_opt ~pos regexp s with
    | Some group ->
        muls s ~pos:(Re.Group.stop group 0) ~acc:(acc + product group)
    | None -> acc
  in

  List.fold_left (fun acc it -> acc + muls it) 0 input

let part2 input =
  let regexp =
    Re.Pcre.re {|mul\((\d+),(\d+)\)|don't\(\)|do\(\)|} |> Re.compile
  in
  let rec muls ?(pos = 0) acc s =
    match Re.exec_opt ~pos regexp s with
    | Some group ->
        let sum, enabled = acc in
        let new_acc =
          match Re.Group.get group 0 with
          | "do()" -> (sum, true)
          | "don't()" -> (sum, false)
          | _ -> if enabled then (sum + product group, enabled) else acc
        in
        muls ~pos:(Re.Group.stop group 0) new_acc s
    | None -> acc
  in

  List.fold_left (fun acc it -> muls acc it) (0, true) input |> fst
