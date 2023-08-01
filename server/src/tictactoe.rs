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

    fn row_winner(&self, board: Board) -> Option<State> {
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
                Some(State::from_i32(cell).unwrap());
            }
        }

        None
    }

    fn column_winner(&self, board: Board) -> Option<State> {
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
                Some(State::from_i32(current_cell).unwrap());
            }
        }

        None
    }

    fn diagonal_winner(&self, board: Board) -> Option<State> {
        for (r, row) in board.rows.iter().enumerate() {
            let mut is_same_slash = false;
            let mut is_same_backslash = false;
            let current_cell_slash = if row.cells[r].state != State::Empty as i32 {
                Some(row.cells[r].state)
            } else {
                None
            };
            let current_cell_backslash =
                if row.cells[row.cells.len() - 1 - r].state != State::Empty as i32 {
                    Some(row.cells[row.cells.len() - 1 - r].state)
                } else {
                    None
                };

            for cell in row.cells.iter() {
                if let Some(slash) = current_cell_slash {
                    if cell.state == slash {
                        is_same_slash = true;
                    }
                }
                if let Some(backslash) = current_cell_backslash {
                    if cell.state == backslash {
                        is_same_backslash = true;
                    }
                }
            }

            if is_same_slash {
                return Some(State::from_i32(current_cell_slash.unwrap()).unwrap());
            }

            if is_same_backslash {
                return Some(State::from_i32(current_cell_backslash.unwrap()).unwrap());
            }
        }

        None
    }

    pub fn winner(&self, board: Board) -> Option<State> {
        if let Some(winner) = self.row_winner(board.clone()) {
            return Some(winner);
        }

        if let Some(winner) = self.column_winner(board.clone()) {
            return Some(winner);
        }

        if let Some(winner) = self.diagonal_winner(board.clone()) {
            return Some(winner);
        }

        None
    }

    pub fn is_terminal(&self, board: Board) -> bool {
        if let Some(_) = self.winner(board.clone()) {
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

    fn utility(&self, board: Board) -> i32 {
        match self.winner(board.clone()) {
            Some(winner) => match winner {
                State::X => 1,
                State::O => -1,
                State::Empty => 0,
            },
            None => 0,
        }
    }

    pub fn minimax(&self, board: Board) -> Option<Action> {
        match self.turn(board.clone()) {
            State::X => Some(self.max_value(board.clone())),
            State::O => Some(self.min_value(board.clone())),
            State::Empty => None,
        };

        None
    }

    fn max_value(&self, board: Board) -> (i32, Option<Action>) {
        if self.is_terminal(board.clone()) {
            return (self.utility(board.clone()), None);
        }

        self.possible_actions(board.clone())
            .actions
            .iter()
            .map(|a| (self.min_value(self.result(board.clone(), a.clone()))))
            .max_by_key(|(v, _)| v.clone())
            .unwrap()
    }

    fn min_value(&self, board: Board) -> (i32, Option<Action>) {
        if self.is_terminal(board.clone()) {
            return (self.utility(board.clone()), None);
        }

        self.possible_actions(board.clone())
            .actions
            .iter()
            .map(|a| (self.max_value(self.result(board.clone(), a.clone()))))
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

        assert_eq!(tic_tac_toe.winner(result), None);
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

        assert_eq!(tic_tac_toe.winner(board), Some(State::X));

        let mut board = tic_tac_toe.initial_state();
        board.rows[1].cells[0] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[1].cells[2] = cell_x.clone();

        assert_eq!(tic_tac_toe.winner(board), Some(State::X));

        let mut board = tic_tac_toe.initial_state();
        board.rows[2].cells[0] = cell_x.clone();
        board.rows[2].cells[1] = cell_x.clone();
        board.rows[2].cells[2] = cell_x.clone();

        assert_eq!(tic_tac_toe.winner(board), Some(State::X));
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
        assert_eq!(tic_tac_toe.winner(board), Some(State::X));

        let mut board = tic_tac_toe.initial_state();
        board.rows[0].cells[1] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[2].cells[1] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), Some(State::X));

        let mut board = tic_tac_toe.initial_state();
        board.rows[0].cells[2] = cell_x.clone();
        board.rows[1].cells[2] = cell_x.clone();
        board.rows[2].cells[2] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), Some(State::X));
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
        assert_eq!(tic_tac_toe.winner(board), Some(State::X));

        let mut board = tic_tac_toe.initial_state();
        board.rows[0].cells[2] = cell_x.clone();
        board.rows[1].cells[1] = cell_x.clone();
        board.rows[2].cells[0] = cell_x.clone();
        assert_eq!(tic_tac_toe.winner(board), Some(State::X));
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
