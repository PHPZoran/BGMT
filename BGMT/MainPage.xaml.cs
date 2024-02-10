using CommunityToolkit.Maui.Storage;
using System.Runtime.CompilerServices;

namespace BGMT
{
    public partial class MainPage : ContentPage
    {
        private string projectFolderPath;
        public MainPage()
        {
            InitializeComponent();

        }

        // File Subitem Clicked Functionality
        private async void NavBarFile_New(object sender, EventArgs e) 
        {
            try 
            {
                var folder = await FolderPicker.PickAsync(default);
                this.projectFolderPath = folder.Folder.Path;
                target.Text = this.projectFolderPath;

                this.menuItemImport.IsEnabled = true;
                this.menuItemExport.IsEnabled = true;
            }
            catch(Exception ex) { Console.WriteLine(ex.ToString()); }
        }
        void NavBarFile_Open(object sender, EventArgs e) 
        { 
            target.Text = "Open Pressed";

        }
        void NavBarFile_Import(object sender, EventArgs e) { target.Text = "Import Pressed"; }
        void NavBarFile_Export(object sender, EventArgs e) { target.Text = "Export Pressed"; }
        void NavBarFile_Exit(object sender, EventArgs e) { target.Text = "Exit Pressed"; }

        // Help Subitem Clicked Functionality
        void NavBarHelp_Settings(object sender, EventArgs e) { target.Text = "Settings Pressed"; }
        void NavBarHelp_Help(object sender, EventArgs e) { target.Text = "Help Pressed"; }
    }

}
