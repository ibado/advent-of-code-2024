let part1 input =
  let regexp = Re.Pcre.re {|mul\((\d+),(\d+)\)|} |> Re.compile in
  let rec muls_aux s pos acc =
    match Re.exec_opt ~pos regexp s with
    | Some groups ->
        let sum =
          int_of_string (Re.Group.get groups 1)
          * int_of_string (Re.Group.get groups 2)
        in
        muls_aux s (Re.Group.stop groups 0) (acc + sum)
    | None -> acc
  in
  let muls s = muls_aux s 0 0 in

  List.fold_left (fun acc it -> acc + muls it) 0 input

let part2 _input = 0
