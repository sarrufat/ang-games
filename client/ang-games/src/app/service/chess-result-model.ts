export interface ChessResult {
    done: boolean;
    ms: number;
    combinations: Combination[];
}

export interface Combination {
    positions: Position[];
}

export interface Position {
    piece: string;
    x: number;
    y: number;
}

export interface TaskId {
    taskId: string;
}