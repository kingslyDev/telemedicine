import React from 'react';

export const Footer = () => {
  return (
    <div>
      {' '}
      {/* Footer */}
      <footer className="pt-60">
        <div className="footer bg-gradient-to-r p-20 from-green-700 to-green-300">
          <ul className="flex flex-col md:flex-row justify-between">
            <li className="mb-10 md:mb-0">
              <div className="logoFooter flex gap-3">
                {/* <img src={logoFooter} alt="MindTrack Logo" /> */}
                <h1 className="text-white font-semibold text-3xl mt-1">MindTrack</h1>
              </div>
            </li>
            <li className="mb-10 md:mb-0">
              <div className="textService mt-2">
                <ul>
                  <li>
                    <h1 className="text-white font-bold text-lg underline pb-5">Service</h1>
                  </li>
                  <li>
                    <h1 className="text-white underline">Tes Kesehatan Mental</h1>
                  </li>
                  <li>
                    <h1 className="text-white underline">Konseling Mental</h1>
                  </li>
                  <li>
                    <h1 className="text-white underline">Janji Dokter</h1>
                  </li>
                </ul>
              </div>
            </li>
            <li className="mb-10 md:mb-0">
              <div className="textContact mt-2">
                <ul>
                  <li>
                    <h1 className="text-white font-bold text-lg underline pb-5">Contact</h1>
                  </li>
                  <li className="flex gap-4">
                    {/* <img src={call} alt="Phone Icon" /> */}
                    <h1 className="text-white underline">+62 823-3133-3221</h1>
                  </li>
                  <li className="flex gap-4">
                    {/* <img src={message} alt="Message Icon" /> */}
                    <h1 className="text-white underline">mindtrack@domain.com</h1>
                  </li>
                  <li className="flex gap-4">
                    {/* <img src={location} alt="Location Icon" /> */}
                    <h1 className="text-white underline">Politeknik Elektronika Negeri Surabaya</h1>
                  </li>
                </ul>
              </div>
            </li>
            <li>
              <div className="textLinks mt-2">
                <ul>
                  <li>
                    <h1 className="text-white font-bold text-lg underline pb-5">Links</h1>
                  </li>
                  <li>
                    <h1 className="text-white underline">Kebijakan Privasi</h1>
                  </li>
                  <li>
                    <h1 className="text-white underline">Ketentuan Penggunaan</h1>
                  </li>
                </ul>
              </div>
            </li>
          </ul>
          <hr className="mt-20 border-2" />
          <h1 className="text-white text-center mt-10 underline">Copyright 2024 @mindtrack all rights reserved</h1>
        </div>
      </footer>
    </div>
  );
};
