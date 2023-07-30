use tonic::{transport::Server, Request, Response, Status};

use tictactoegrpc::tic_tac_toe_grpc_server::{TicTacToeGrpc, TicTacToeGrpcServer};
use tictactoegrpc::{Actions, Board, BoardWithAction, Player, State, Terminal};

use tictactoe::TicTacToe;
mod tictactoe;

pub mod tictactoegrpc {
    tonic::include_proto!("tictactoe");
}

#[derive(Default)]
struct TicTacToeService {
    tic_tac_toe: TicTacToe,
}

#[tonic::async_trait]
impl TicTacToeGrpc for TicTacToeService {
    async fn initial_state(&self, _: Request<()>) -> std::result::Result<Response<Board>, Status> {
        Ok(Response::new(self.tic_tac_toe.initial_state()))
    }

    async fn turn(&self, request: Request<Board>) -> Result<Response<Player>, Status> {
        let board = request.into_inner();
        Ok(Response::new(Player {
            player: self.tic_tac_toe.turn(board) as i32,
        }))
    }

    async fn possible_actions(
        &self,
        request: Request<Board>,
    ) -> std::result::Result<Response<Actions>, Status> {
        let board = request.into_inner();
        Ok(Response::new(self.tic_tac_toe.possible_actions(board)))
    }

    async fn result(
        &self,
        request: Request<BoardWithAction>,
    ) -> std::result::Result<Response<Board>, Status> {
        let request = request.into_inner();
        let board = request.board.unwrap();
        let action = request.action.unwrap();
        Ok(Response::new(self.tic_tac_toe.result(board, action)))
    }

    async fn winner(&self, request: Request<Board>) -> Result<Response<Player>, Status> {
        let board = request.into_inner();
        let player = if let Some(player) = self.tic_tac_toe.winner(board) {
            player as i32
        } else {
            State::Empty as i32
        };
        Ok(Response::new(Player { player }))
    }

    async fn is_terminal(&self, request: Request<Board>) -> Result<Response<Terminal>, Status> {
        let board = request.into_inner();
        Ok(Response::new(Terminal {
            terminal: self.tic_tac_toe.is_terminal(board),
        }))
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
