import { useQuery } from '@tanstack/react-query';
import './PlayerList.css';
import { fetchPlayers } from './players';

function PlayerList() {
    const { data: players, isPending, isError, error } = useQuery({
        queryKey: ['players'],
        queryFn: fetchPlayers,
    });

    if (isPending) {
        return <div className="player-container status-center">Loading players…</div>;
    }

    if (isError) {
        return <div className="player-container status-center">Failed to load players: {error.message}</div>;
    }

    return (
        <div className="player-container">
            {players.map((player, index) => (
                <div key={player.id} className="row">
                    <div>
                        <img
                            className="avatar"
                            src={player.avatar}
                            alt={player.name}
                            width={48}  /* 👈 Add this */
                            height={48} /* 👈 Add this */
                            loading={index < 5 ? 'eager' : 'lazy'}
                            fetchPriority={index < 5 ? 'high' : 'auto'}
                        />
                    </div>

                    <div className="col">
                        <span className="name">{player.name}</span>
                        <span className="age">Age {player.age}</span>
                    </div>
                </div>
            ))}
        </div>
    );
}

export default PlayerList;