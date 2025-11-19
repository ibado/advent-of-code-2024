let part1 input =
  let regexp = Re.Pcre.re {|mul\((\d+),(\d+)\)|} |> Re.compile in
  let rec muls ?(pos = 0) ?(acc = 0) s =
    match Re.exec_opt ~pos regexp s with
    | Some groups ->
        let sum =
          int_of_string (Re.Group.get groups 1)
          * int_of_string (Re.Group.get groups 2)
        in
        muls s ~pos:(Re.Group.stop groups 0) ~acc:(acc + sum)
    | None -> acc
  in

  List.fold_left (fun acc it -> acc + muls it) 0 input

let part2 _input = 0
