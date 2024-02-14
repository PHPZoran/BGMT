namespace TemplateBGMT.UI_Elements;

public partial class FolderDirectory : ContentView
{
    private string path;
	public FolderDirectory(string LoadOption)
	{
		InitializeComponent();
	}

    public void AddFileFolderDirectory(object sender, EventArgs e)
    {

    }

    public void LoadFileFolderDirectory(object sender, EventArgs e)
    {

    }

    public void DeleteFileFolderDirectory(object sender, EventArgs e)
    {

    }

    public ListView GetFolderView() { return this.lv_FolderDirectory; }
}