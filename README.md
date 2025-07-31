# TMDB-CLI-Tool

#  TMDB CLI App

A command-line tool built with Go to fetch movie data from [The Movie Database (TMDB)](https://www.themoviedb.org/). You can list **popular**, **top-rated**, or **upcoming** movies straight from your terminal.

---

##  Features

- Fetches data from TMDB v3 API
- Supports listing:
  -  Popular Movies
  -  Top Rated Movies
  -  Upcoming Movies
- Secure authentication using Bearer Token via `.env` file
- Clean CLI experience

---

##  Requirements

- [Go](https://golang.org/) 1.18 or higher
- A [TMDB API Bearer Token](https://developer.themoviedb.org/docs/getting-started)

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/tmdb-cli-app.git
cd tmdb-cli-app
```

### 2. Create .env File
Create a .env file in the root directory:

```bash
TMDB_API_TOKEN=your_tmdb_bearer_token_here
```

Make sure the token does not contain quotes or extra whitespace. The program will automatically TrimSpace() and remove any accidental "

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the App
```bash
go run main.go --type [playing | popular | top | upcoming]
```
### Example:
```bash
go run main.go --type popular
```
### Output:

Title: War of the Worlds
Release Date: 2025-07-29
Overview: Will Radford is a top cyber-security analyst for Homeland Security who tracks potential threats to national security through a mass surveillance program, until one day an attack by an unknown entity leads him to question whether the government is hiding something from him... and from the rest of the world.
Vote Average: 4.8

Title: How to Train Your Dragon
Release Date: 2025-06-06
Overview: On the rugged isle of Berk, where Vikings and dragons have been bitter enemies for generations, Hiccup stands apart, defying centuries of tradition when he befriends Toothless, a feared Night Fury dragon. Their unlikely bond reveals the true nature of dragons, challenging the very foundations of Viking society.
Vote Average: 8.1

Title: Demon Slayer: Kimetsu no Yaiba Infinity Castle
Release Date: 2025-07-18
Overview: As the Demon Slayer Corps members and Hashira engaged in a group strength training program, the Hashira Training, in preparation for the forthcoming battle against the demons, Muzan Kibutsuji appears at the Ubuyashiki Mansion. With the head of the Demon Corps in danger, Tanjiro and the Hashira rush to the headquarters but are plunged into a deep descent to a mysterious space by the hands of Muzan Kibutsuji.  The destination of where Tanjiro and Demon Slayer Corps have fallen is the demons' stronghold – the Infinity Castle. And so, the battleground is set as the final battle between the Demon Slayer Corps and the demons ignites.
Vote Average: 7.0

Title: Lilo & Stitch
Release Date: 2025-05-17
Overview: The wildly funny and touching story of a lonely Hawaiian girl and the fugitive alien who helps to mend her broken family.
Vote Average: 7.3

Title: M3GAN 2.0
Release Date: 2025-06-25
Overview: After the underlying tech for M3GAN is stolen and misused by a powerful defense contractor to create a military-grade weapon known as Amelia, M3GAN's creator Gemma realizes that the only option is to resurrect M3GAN and give her a few upgrades, making her faster, stronger, and more lethal.
Vote Average: 7.6

Title: Happy Gilmore 2
Release Date: 2025-07-25
Overview: Happy Gilmore isn't done with golf — not by a long shot. Since his retirement after his first Tour Championship win, Gilmore returns to finance his daughter's ballet classes.
Vote Average: 6.8

Title: The Fantastic 4: First Steps
Release Date: 2025-07-23
Overview: Against the vibrant backdrop of a 1960s-inspired, retro-futuristic world, Marvel's First Family is forced to balance their roles as heroes with the strength of their family bond, while defending Earth from a ravenous space god called Galactus and his enigmatic Herald, Silver Surfer.
Vote Average: 7.4

Title: Bride Hard
Release Date: 2025-06-19
Overview: Sam is a secret agent whose toughest mission to date is pleasing her bride-to-be best friend at a lavish destination wedding. When a team of mercenaries crashes the party and takes the guests hostage, Sam is thrown into a fight unlike any before — one where she can’t risk blowing her cover or ruining the big day. As she takes on the bad guys in a high-stakes battle disguised as a fairy-tale affair, she realizes the real threat might be closer than she thinks.
Vote Average: 6.2

Title: Superman
Release Date: 2025-07-09
Overview: Superman, a journalist in Metropolis, embarks on a journey to reconcile his Kryptonian heritage with his human upbringing as Clark Kent.
Vote Average: 7.4

Title: Ice Road: Vengeance
Release Date: 2025-06-27
Overview: Big rig ice road driver Mike McCann travels to Nepal to scatter his late brother’s ashes on Mt. Everest. While on a packed tour bus traversing the deadly 12,000 ft. terrain of the infamous Road to the Sky, McCann and his mountain guide encounter a group of mercenaries and must fight to save themselves, the busload of innocent travelers, and the local villagers’ homeland.
Vote Average: 6.8

Title: Man with No Past
Release Date: 2025-01-13
Overview: Waking up in an unfamiliar city, a man with no memory must confront the mysteries of his own identity. However, his desperate search to uncover the past pits him against a powerful enemy, leading to a showdown that ultimately reveals the truth.
Vote Average: 6.7

Title: Karate Kid: Legends
Release Date: 2025-05-08
Overview: After a family tragedy, kung fu prodigy Li Fong is uprooted from his home in Beijing and forced to move to New York City with his mother. When a new friend needs his help, Li enters a karate competition – but his skills alone aren't enough. Li's kung fu teacher Mr. Han enlists original Karate Kid Daniel LaRusso for help, and Li learns a new way to fight, merging their two styles into one for the ultimate martial arts showdown.
Vote Average: 7.2

Title: Jurassic World Rebirth
Release Date: 2025-07-01
Overview: Five years after the events of Jurassic World Dominion, covert operations expert Zora Bennett is contracted to lead a skilled team on a top-secret mission to secure genetic material from the world's three most massive dinosaurs. When Zora's operation intersects with a civilian family whose boating expedition was capsized, they all find themselves stranded on an island where they come face-to-face with a sinister, shocking discovery that's been hidden from the world for decades.
Vote Average: 6.3

Title: Ballerina
Release Date: 2025-06-04
Overview: Taking place during the events of John Wick: Chapter 3 – Parabellum, Eve Macarro begins her training in the assassin traditions of the Ruska Roma.
Vote Average: 7.5

Title: KPop Demon Hunters
Release Date: 2025-06-20
Overview: When K-pop superstars Rumi, Mira and Zoey aren't selling out stadiums, they're using their secret powers to protect their fans from supernatural threats.
Vote Average: 8.5

Title: Thunderbolts*
Release Date: 2025-04-30
Overview: After finding themselves ensnared in a death trap, seven disillusioned castoffs must embark on a dangerous mission that will force them to confront the darkest corners of their pasts.
Vote Average: 7.4

Title: Final Destination Bloodlines
Release Date: 2025-05-14
Overview: Plagued by a violent recurring nightmare, college student Stefanie heads home to track down the one person who might be able to break the cycle and save her family from the grisly demise that inevitably awaits them all.
Vote Average: 7.2

Title: Almost Cops
Release Date: 2025-07-10
Overview: When an overeager community officer and a reckless ex-detective are forced to team up, plenty of chaos ensues on the streets of Rotterdam.
Vote Average: 5.9

Title: Dangerous Animals
Release Date: 2025-06-05
Overview: A savvy and free-spirited surfer is abducted by a shark-obsessed serial killer. Held captive on his boat, she must figure out how to escape before he carries out a ritualistic feeding to the sharks below.
Vote Average: 6.5

Title: Angels Fallen: Warriors of Peace
Release Date: 2024-07-09
Overview: When an Iraq War veteran receives a calling from a higher power, he embarks on a mission to stop a fallen angel from raising an army of the dead to take over the world.
Vote Average: 5.9

Movies fetched successfully.

How Authentication Works
Uses Bearer Token provided by TMDB
Loaded securely from .env using github.com/joho/godotenv
```bash
import (
    "os"
    "strings"
)

token := strings.TrimSpace(os.Getenv("TMDB_API_TOKEN"))
token = strings.Trim(token, "\"")
```



