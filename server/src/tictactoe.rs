use crate::tictactoegrpc::{Action, Actions, Board, Cell, Row, State};

#[derive(Default)]
pub struct TicTacToe {}

impl TicTacToe {
    pub fn initial_state(&self) -> Board {
        Board {
            rows: vec![
                Row {
                    cells: vec![
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                    ],
                },
                Row {
                    cells: vec![
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                    ],
                },
                Row {
                    cells: vec![
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                    ],
                },
            ],
        }
    }

    pub fn turn(&self, board: &Board) -> State {
        let count_x: usize = board
            .rows
            .iter()
            .map(|row| {
                row.cells
                    .iter()
                    .filter(|&cell| cell.state == State::X as i32)
                    .count()
            })
            .sum();

        let count_o: usize = board
            .rows
            .iter()
            .map(|row| {
                row.cells
                    .iter()
                    .filter(|&cell| cell.state == State::O as i32)
                    .count()
            })
            .sum();

        if count_x == count_o {
            State::X
        } else if count_x > count_o {
            State::O
        } else {
            panic!("Invalid board");
        }
    }

    pub fn possible_actions(&self, board: &Board) -> Actions {
        let mut actions = Actions::default();

        for (r, row) in board.rows.iter().enumerate() {
            for (c, cell) in row.cells.iter().enumerate() {
                if cell.state == State::Empty as i32 {
                    actions.actions.push(Action {
                        row: r as i32,
                        column: c as i32,
                    });
                }
            }
        }

        actions
    }

    pub fn result(&self, mut board: Board, action: &Action) -> Result<Board, &str> {
        let player = self.turn(&board);

        if board.rows[action.row as usize].cells[action.column as usize]
            == (Cell {
                state: State::Empty as i32,
            })
        {
            board.rows[action.row as usize].cells[action.column as usize] = Cell {
                state: player as i32,
            };

            return Ok(board);
        }

        Err("Invalid action")
    }

    fn row_winner(&self, board: &Board) -> Option<State> {
        for row in board.rows.iter() {
            let xs = row
                .cells
                .iter()
                .filter(|c| c.state == State::X as i32)
                .count();
            if xs == 3 {
                return Some(State::X);
            }
            let os = row
                .cells
                .iter()
                .filter(|c| c.state == State::O as i32)
                .count();
            if os == 3 {
                return Some(State::O);
            }
        }

        None
    }

    fn column_winner(&self, board: &Board) -> Option<State> {
        for i in 0..3 {
            if board.rows[0].cells[i].state != State::Empty as i32
                && board.rows[0].cells[i].state == board.rows[1].cells[i].state
                && board.rows[1].cells[i].state == board.rows[2].cells[i].state
            {
                return Some(State::from_i32(board.rows[0].cells[i].state).unwrap());
            }
        }

        None
    }

    fn diagonal_winner(&self, board: &Board) -> Option<State> {
        if board.rows[0].cells[0].state != State::Empty as i32
            && board.rows[0].cells[0].state == board.rows[1].cells[1].state
            && board.rows[1].cells[1].state == board.rows[2].cells[2].state
        {
            return Some(State::from_i32(board.rows[0].cells[0].state).unwrap());
        }
        if board.rows[0].cells[2].state != State::Empty as i32
            && board.rows[0].cells[2].state == board.rows[1].cells[1].state
            && board.rows[1].cells[1].state == board.rows[2].cells[0].state
        {
            return Some(State::from_i32(board.rows[0].cells[2].state).unwrap());
        }

        None
    }

    pub fn winner(&self, board: &Board) -> Option<State> {
        if let Some(winner) = self.row_winner(&board) {
            return Some(winner);
        }

        if let Some(winner) = self.column_winner(&board) {
            return Some(winner);
        }

        if let Some(winner) = self.diagonal_winner(&board) {
            return Some(winner);
        }

        None
    }

    pub fn is_terminal(&self, board: &Board) -> bool {
        if let Some(_) = self.winner(&board) {
            return true;
        }
        let empty_cell = Cell {
            state: State::Empty as i32,
        };

        let is_any = board.rows.iter().any(|row| row.cells.contains(&empty_cell));

        !is_any
    }

    fn utility(&self, board: &Board) -> Option<i32> {
        match self.winner(&board) {
            Some(winner) => match winner {
                State::X => Some(1),
                State::O => Some(-1),
                State::Empty => Some(0),
            },
            None => Some(0),
        }
    }

    pub fn minimax(&self, board: &Board) -> Option<Action> {
        if self.is_terminal(&board) {
            return None;
        }

        if self.turn(&board) == State::X {
            return self.max_value(&board).1;
        } else {
            return self.min_value(&board).1;
        }
    }

    fn max_value(&self, board: &Board) -> (i32, Option<Action>) {
        if self.is_terminal(&board) {
            if let Some(utility) = self.utility(&board) {
                return (utility, None);
            } else {
                panic!("utility is None")
            }
        }

        self.possible_actions(&board)
            .actions
            .iter()
            .map(|a| {
                let result = self.result(board.clone(), &a);
                self.min_value(&result.unwrap())
            })
            .max_by_key(|(v, _)| v.clone())
            .unwrap()
    }

    fn min_value(&self, board: &Board) -> (i32, Option<Action>) {
        if self.is_terminal(&board) {
            if let Some(utility) = self.utility(&board) {
                return (utility, None);
            } else {
                panic!("utility is None")
            }
        }

        self.possible_actions(&board)
            .actions
            .iter()
            .map(|a| {
                let result = self.result(board.clone(), &a);
                self.max_value(&result.unwrap())
            })
            .min_by_key(|(v, _)| v.clone())
            .unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_initial_state() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();

        assert_eq!(board.rows.len(), 3);
        assert_eq!(board.rows[0].cells.len(), 3);
        assert_eq!(board.rows[1].cells.len(), 3);
        assert_eq!(board.rows[2].cells.len(), 3);

        let empty_row = Row {
            cells: vec![
                Cell {
                    state: State::Empty as i32,
                },
                Cell {
                    state: State::Empty as i32,
                },
                Cell {
                    state: State::Empty as i32,
                },
            ],
        };

        for row in board.rows.iter() {
            assert_eq!(row, &empty_row);
        }
    }

    #[test]
    fn test_turn() {
        let tic_tac_toe = TicTacToe::default();
        let mut board = tic_tac_toe.initial_state();
        assert_eq!(tic_tac_toe.turn(&board), State::X);

        board.rows[0].cells[0] = Cell {
            state: State::X as i32,
        };
        assert_eq!(tic_tac_toe.turn(&board), State::O);

        board.rows[0].cells[1] = Cell {
            state: State::O as i32,
        };
        assert_eq!(tic_tac_toe.turn(&board), State::X);

        board.rows[1].cells[0] = Cell {
            state: State::X as i32,
        };
        assert_eq!(tic_tac_toe.turn(&board), State::O);

        board.rows[1].cells[1] = Cell {
            state: State::O as i32,
        };
        assert_eq!(tic_tac_toe.turn(&board), State::X);
    }

    #[test]
    fn test_possible_actions() {
        let tic_tac_toe = TicTacToe::default();
        let mut board = tic_tac_toe.initial_state();

        let actions = tic_tac_toe.possible_actions(&board);
        assert_eq!(actions.actions.len(), 9);

        board.rows[0].cells[0] = Cell {
            state: State::X as i32,
        };
        board.rows[0].cells[1] = Cell {
            state: State::O as i32,
        };
        let actions = tic_tac_toe.possible_actions(&board);
        assert_eq!(actions.actions.len(), 7);
    }

    #[test]
    fn test_result() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();

        let action = Action { row: 0, column: 1 };
        let board = tic_tac_toe.result(board, &action).unwrap();

        assert_eq!(board.rows[0].cells[1].state, State::X as i32);

        let action = Action { row: 2, column: 2 };
        let result = tic_tac_toe.result(board, &action);

        assert_eq!(result.unwrap().rows[2].cells[2].state, State::O as i32);
    }

    #[test]
    fn test_winner() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();

        // . X .
        // . . .
        // . . .
        let action = Action { row: 0, column: 1 };
        let board = tic_tac_toe.result(board, &action).unwrap();
        assert_eq!(board.rows[0].cells[1].state, State::X as i32);
        assert_eq!(tic_tac_toe.winner(&board), None);

        // O X .
        // . . .
        // . . .
        let action = Action { row: 0, column: 0 };
        let board = tic_tac_toe.result(board, &action).unwrap();
        assert_eq!(board.rows[0].cells[0].state, State::O as i32);
        assert_eq!(tic_tac_toe.winner(&board), None);

        // O X .
        // . X .
        // . . .
        let action = Action { row: 1, column: 1 };
        let board = tic_tac_toe.result(board, &action).unwrap();
        assert_eq!(board.rows[1].cells[1].state, State::X as i32);
        assert_eq!(tic_tac_toe.winner(&board), None);

        // O X .
        // O X .
        // . . .
        let action = Action { row: 1, column: 0 };
        let board = tic_tac_toe.result(board, &action).unwrap();
        assert_eq!(board.rows[1].cells[0].state, State::O as i32);
        assert_eq!(tic_tac_toe.winner(&board), None);

        // O X .
        // O X .
        // . X .
        let action = Action { row: 2, column: 1 };
        let board = tic_tac_toe.result(board, &action).unwrap();
        assert_eq!(board.rows[2].cells[1].state, State::X as i32);
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));
    }

    #[test]
    fn test_row_winner() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 2 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));

        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 2 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));

        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 2 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));
    }

    #[test]
    fn test_column_winner() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 0 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));

        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 1 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));

        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 2 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 2 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 2 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));
    }

    #[test]
    fn test_slash_winner() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 2 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));

        let board = tic_tac_toe.initial_state();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 2 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 0, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 1 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 1, column: 0 })
            .unwrap();
        let board = tic_tac_toe
            .result(board, &Action { row: 2, column: 0 })
            .unwrap();
        assert_eq!(tic_tac_toe.winner(&board), Some(State::X));
    }

    #[test]
    fn test_is_terminal() {
        let board = Board {
            rows: vec![
                Row {
                    cells: vec![
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::O as i32,
                        },
                        Cell {
                            state: State::X as i32,
                        },
                    ],
                },
                Row {
                    cells: vec![
                        Cell {
                            state: State::O as i32,
                        },
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::O as i32,
                        },
                    ],
                },
                Row {
                    cells: vec![
                        Cell {
                            state: State::O as i32,
                        },
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::O as i32,
                        },
                    ],
                },
            ],
        };

        let tic_tac_toe = TicTacToe::default();
        assert_eq!(tic_tac_toe.winner(&board), None);
        assert_eq!(tic_tac_toe.is_terminal(&board), true);
    }

    #[test]
    fn test_minimax() {
        let tic_tac_toe = TicTacToe::default();
        let board = Board {
            rows: vec![
                Row {
                    cells: vec![
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::O as i32,
                        },
                    ],
                },
                Row {
                    cells: vec![
                        Cell {
                            state: State::O as i32,
                        },
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::O as i32,
                        },
                    ],
                },
                Row {
                    cells: vec![
                        Cell {
                            state: State::X as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::Empty as i32,
                        },
                    ],
                },
            ],
        };
        // there is no winner
        assert_eq!(tic_tac_toe.minimax(&board), None);
    }
}
