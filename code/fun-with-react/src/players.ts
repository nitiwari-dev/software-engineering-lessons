export type Player = {
  id: number;
  name: string;
  age: number;
  avatar: string;
};

// Shape of the relevant slice of the randomuser.me response.
// Only fields we actually use are declared.
type RandomUserResponse = {
  results: Array<{
    name: { first: string; last: string };
    dob: { age: number };
    picture: { large: string };
  }>;
};

export async function fetchPlayers(): Promise<Player[]> {
  const response = await fetch('https://randomuser.me/api/?results=100&nat=in');

  if (!response.ok) {
    throw new Error(`Failed to fetch players: ${response.status} ${response.statusText}`);
  }

  const data: RandomUserResponse = await response.json();

  return data.results.map((user, i) => ({
    id: i + 1,
    name: `${user.name.first} ${user.name.last}`,
    age: user.dob.age,
    avatar: user.picture.large,
  }));
}