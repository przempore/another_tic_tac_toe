use tonic::{transport::Server, Request, Response, Status};

use tictactoe::tic_tac_toe_grpc_server::{TicTacToeGrpc, TicTacToeGrpcServer};
use tictactoe::{Action, Actions, Board, BoardWithAction, Cell, Player, Row, State};

pub mod tictactoe {
    tonic::include_proto!("tictactoe");
}

#[derive(Default)]
struct TicTacToeService {} // TODO: Service should contain TicTacToe object

#[tonic::async_trait]
impl TicTacToeGrpc for TicTacToeService {
    async fn initial_state(&self, _: Request<()>) -> std::result::Result<Response<Board>, Status> {
        let board = Board {
            rows: vec![
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
                Row {
                    cells: vec![
                        Cell {
                            state: State::Empty as i32,
                        },
                        Cell {
                            state: State::X as i32,
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
                            state: State::X as i32,
                        },
                    ],
                },
            ],
        };

        Ok(Response::new(board))
    }

    async fn turn(&self, request: Request<Board>) -> Result<Response<Player>, Status> {
        let board = request.get_ref();

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

        return Ok(Response::new(Player {
            player: player as i32,
        }));
    }

    async fn possible_actions(
        &self,
        request: Request<Board>,
    ) -> std::result::Result<Response<Actions>, Status> {
        let board = request.get_ref();
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

        Ok(Response::new(actions))
    }

    async fn result(
        &self,
        request: Request<BoardWithAction>,
    ) -> std::result::Result<Response<Board>, Status> {
        let request = request.into_inner();
        let mut board = request.board.unwrap();
        let action = request.action.unwrap();

        let player = self.turn(Request::new(board.clone())).await?;

        board.rows[action.row as usize].cells[action.column as usize] = Cell {
            state: player.get_ref().player as i32,
        };

        Ok(Response::new(Board::default()))
    }
}

#[derive(Default)]
struct TicTacToe {}

impl TicTacToe {
    fn turn(&self, board: Board) -> State {
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
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse().unwrap();
    let tic_tac_toe_service = TicTacToeService::default();

    println!("tic_tac_toe listening on {}", addr);

    Server::builder()
        .add_service(TicTacToeGrpcServer::new(tic_tac_toe_service))
        .serve(addr)
        .await?;

    Ok(())
}
