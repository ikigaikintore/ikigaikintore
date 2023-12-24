import * as WeatherCard from '@/src/cards/weather'

export default function Page() {
    return (
        <section className="mt-12 mx-auto px-4 max-w-screen-xl md:px-8">
            <header className="mainLinks">
                My links here
            </header>
            <main className="lower-part grid grid-cols-2 gap-4">
                <aside className="linksPages">
                    Link elements
                </aside>
                <section className="bodyCards grid grid-cols-4 gap-4">
                    <div>
                        <WeatherCard.Component/>
                    </div>
                </section>
            </main>
        </section>
    )
}