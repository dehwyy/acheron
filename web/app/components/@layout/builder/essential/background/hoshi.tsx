"use client"
import "./hoshi.scss"

export function Hoshi({ children }: { children: React.ReactNode }) {
    return (
        <>

            <div
                className="stars__wrapper"
            >
                <div className="stars" />
                <div className="stars2" />
                <div className="stars3" />
            </div>
            {children}
        </>
    )
}
