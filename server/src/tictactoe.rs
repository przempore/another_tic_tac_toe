use tictactoegrpc::{Action, Actions, Board, Cell, Row, State};
pub mod tictactoegrpc {
    tonic::include_proto!("tictactoe");
}

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

    pub fn turn(&self, board: Board) -> State {
        let mut count_x = 0;
        let mut count_y = 0;

        let mut player = State::Empty;

        for row in board.rows.iter() {
            for cell in row.cells.iter() {
                if cell.state == State::X as i32 {
                    count_x += 1;
                } else if cell.state == State::O as i32 {
                    count_y += 1;
                }
            }
        }

        if count_x == count_y {
            player = State::X;
        } else if count_x > count_y {
            player = State::O;
        }

        player
    }

    pub fn possible_actions(&self, board: Board) -> Actions {
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

    pub fn result(&self, mut board: Board, action: Action) -> Board {
        let player = self.turn(board.clone());

        board.rows[action.row as usize].cells[action.column as usize] = Cell {
            state: player as i32,
        };

        board
    }

    fn row_winner(&self, board: Board) -> State {
        let mut winner = State::Empty;

        for row in board.rows.iter() {
            let cell = row.cells[0].state;
            if cell == State::Empty as i32 {
                continue;
            }
            let mut is_same = false;
            for c in row.cells.iter() {
                if c.state != cell {
                    continue;
                }
                is_same = true;
            }
            if is_same {
                winner = State::from_i32(cell).unwrap();
            }
        }

        winner
    }

    fn column_winner(&self, board: Board) -> State {
        let mut winner = State::Empty;
        for (r, row) in board.rows.iter().enumerate() {
            let mut is_same = false;
            let current_cell = row.cells[r].state as i32;
            for (c, _) in row.cells.iter().enumerate() {
                if board.rows[c].cells[r].state == State::Empty as i32 {
                    continue;
                }

                if board.rows[c].cells[r].state == current_cell {
                    is_same = true;
                }
            }

            if is_same {
                winner = State::from_i32(current_cell).unwrap();
            }
        }

        winner
    }

    fn diagonal_winner(&self, board: Board) -> State {
        let mut winner = State::Empty;
        for (r, row) in board.rows.iter().enumerate() {
            let mut is_same_slash = false;
            let mut is_same_backslash = false;
            let current_cell_slash = row.cells[r].state as i32;
            let current_cell_backslash = row.cells[row.cells.len() - 1 - r].state as i32;

            for cell in row.cells.iter() {
                if current_cell_slash == cell.state {
                    is_same_slash = true;
                }
                if current_cell_backslash == cell.state {
                    is_same_backslash = true;
                }
            }

            if is_same_slash {
                winner = State::from_i32(current_cell_slash).unwrap();
            }

            if is_same_backslash {
                winner = State::from_i32(current_cell_backslash).unwrap();
            }
        }

        winner
    }

    pub fn winner(&self, board: Board) -> State {
        let winner = self.row_winner(board.clone());
        if winner != State::Empty {
            return winner;
        }

        let winner = self.column_winner(board.clone());
        if winner != State::Empty {
            return winner;
        }

        let winner = self.diagonal_winner(board.clone());

        winner
    }

    pub fn is_terminal(&self, board: Board) -> bool {
        if self.winner(board.clone()) != State::Empty {
            return true;
        }
        let empty_cell = Cell {
            state: State::Empty as i32,
        };
        for row in board.rows.iter() {
            if row.cells.contains(&empty_cell) {
                return false;
            }
        }

        true
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

        assert_eq!(tic_tac_toe.turn(board.clone()), State::X);

        board.rows[0].cells[0] = Cell {
            state: State::X as i32,
        };

        assert_eq!(tic_tac_toe.turn(board), State::O);
    }

    #[test]
    fn test_possible_actions() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();

        let actions = tic_tac_toe.possible_actions(board);

        assert_eq!(actions.actions.len(), 9);
    }

    #[test]
    fn test_result() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();

        let action = Action { row: 0, column: 1 };
        let board = tic_tac_toe.result(board, action);

        assert_eq!(board.rows[0].cells[1].state, State::X as i32);

        let action = Action { row: 2, column: 2 };
        let result = tic_tac_toe.result(board, action);

        assert_eq!(result.rows[2].cells[2].state, State::O as i32);
    }

    #[test]
    fn test_winner() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();
        let action = Action { row: 0, column: 1 };
        let result = tic_tac_toe.result(board, action);

        assert_eq!(tic_tac_toe.winner(result), State::Empty);
    }

    #[test]
    fn test_row_winner() {
        let tic_tac_toe = TicTacToe::default();
        let mut board = tic_tac_toe.initial_state();
        let cell_x = Cell {
            state: State::X as i32,
        };
        board.rows[0].cells[0] = cell_x.clone();
        board.rows[0].cells[1] = cell_x.clone();
        board.rows[0].cells[2] = cell_x.clone();

        assert_eq!(tic_tac_toe.winner(board), State::X);

        let mut board = tic_tac_toe.initial_state();
        board.rows[1].cells[0] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[1].cells[2] = cell_x.clone();

        assert_eq!(tic_tac_toe.winner(board), State::X);

        let mut board = tic_tac_toe.initial_state();
        board.rows[2].cells[0] = cell_x.clone();
        board.rows[2].cells[1] = cell_x.clone();
        board.rows[2].cells[2] = cell_x.clone();

        assert_eq!(tic_tac_toe.winner(board), State::X);
    }

    #[test]
    fn test_column_winner() {
        let tic_tac_toe = TicTacToe::default();
        let mut board = tic_tac_toe.initial_state();
        let cell_x = Cell {
            state: State::X as i32,
        };

        board.rows[0].cells[0] = cell_x.clone();
        board.rows[1].cells[0] = cell_x.clone();
        board.rows[2].cells[0] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), State::X);

        let mut board = tic_tac_toe.initial_state();
        board.rows[0].cells[1] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[2].cells[1] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), State::X);

        let mut board = tic_tac_toe.initial_state();
        board.rows[0].cells[2] = cell_x.clone();
        board.rows[1].cells[2] = cell_x.clone();
        board.rows[2].cells[2] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), State::X);
    }

    #[test]
    fn test_slash_winner() {
        let tic_tac_toe = TicTacToe::default();
        let mut board = tic_tac_toe.initial_state();
        let cell_x = Cell {
            state: State::X as i32,
        };
        board.rows[0].cells[0] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[2].cells[2] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), State::X);

        let mut board = tic_tac_toe.initial_state();
        board.rows[0].cells[2] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[0].cells[2] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), State::X);
    }

    #[test]
    fn test_is_terminal() {
        let tic_tac_toe = TicTacToe::default();
        let board = tic_tac_toe.initial_state();
        let action = Action { row: 0, column: 1 };
        let result = tic_tac_toe.result(board, action);

        assert_eq!(tic_tac_toe.is_terminal(result), false);
    }
}
